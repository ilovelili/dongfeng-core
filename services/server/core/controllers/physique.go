package controllers

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-error-code"
)

// PhysiqueController physique controller
type PhysiqueController struct {
	repository      *repositories.PhysiqueRepository
	pupilcontroller *PupilController
}

// NewPhysiqueController new physique controller
func NewPhysiqueController() *PhysiqueController {
	return &PhysiqueController{
		repository:      repositories.NewPhysiqueRepository(),
		pupilcontroller: NewPupilController(),
	}
}

// ResolvePhysique resolve physique based on the following steps:
// 1 测量出身高体重后，按照男女性别及年龄，分别比对《0-6岁男、女童体格发育五项指标评价参考值》。
// 2 身高小于P3,怀疑是生长迟缓，根据性别及年龄比对《5岁以下儿童低体重/生长迟缓标准表》。身高小于-2SD，为生长迟缓。
// 3 五项指标评价参考值核对出来后体重小于P10，按照性别及年龄比对《0-6岁按身高测体重》。
// 如身高测体重也小于P10，为营养不良。（年龄测身高和身高测体重两项都小于P3的为重度营养不良；一项小于P10，一项小于P3的或者两项都小于P10的为轻度营养不良。）
// 4 五项指标评价参考值核对出来后体重较重的幼儿
// 五岁以下，按照性别及年龄，核对《5岁以下男/女童身高别体重标准》表，
// 根据身高，大于+1SD为超重，大于+2SD为轻度肥胖，大于+3SD的为中重度肥胖。
// 大于5岁的幼儿，计算BMI指数【体重/身高(米)的平方】，然后比对《5-19岁BMI指数》表，
// 按照性别与年龄比对BMI指数，大于+1SD为超重，大于+2SD为轻度肥胖，大于+3SD的为中重度肥胖。
// 5 5岁以下超重或肥胖的幼儿在计算肥胖度时，根据《5岁以下男/女童身高别体重标准》表，对应相应的身高后，计算公式为实测体重（kg）-中位数/中位数。
func (c *PhysiqueController) ResolvePhysique(physique *models.Physique) (err error) {
	physique.ResolveAge()
	physique.ResolveBMI()

	// step 1. get p zone based on height
	pmasters, err := c.repository.SelectAgeHeightWeightPMasters()
	if err != nil {
		return
	}

	sdmasters, err := c.repository.SelectAgeHeightWeightSDMasters()
	if err != nil {
		return
	}

	hwpmasters, err := c.repository.SelectHeightToWeightPMasters()
	if err != nil {
		return
	}

	hwsdmasters, err := c.repository.SelectHeightToWeightSDMasters()
	if err != nil {
		return
	}

	bmimasters, err := c.repository.SelectBMIMasters()
	if err != nil {
		return
	}

	if found := physique.ResolveAgeHeightP(pmasters); !found {
		err = fmt.Errorf("P height master data not found")
		return
	}

	if found := physique.ResolveAgeWeightP(pmasters); !found {
		err = fmt.Errorf("P weight master data not found")
		return
	}

	if found := physique.ResolveAgeHeightSD(sdmasters); !found {
		// if sd not found... then how about we set it as unknown
		physique.HeightSD = "Unknown"
	}

	if found := physique.ResolveAgeWeightSD(sdmasters); !found {
		// if sd not found... then how about we set it as unknown
		physique.WeightSD = "Unknown"
	}

	if found := physique.ResolveHeightToWeightP(hwpmasters); !found {
		// if hwp not found...
		physique.HeightToWeightP = "Unknown"
	}

	if found := physique.ResolveHeightToWeightSD(hwsdmasters); !found {
		// if hwp not found...
		physique.HeightToWeightSD = "Unknown"
	}

	if found := physique.ResolveFatCofficient(hwsdmasters); !found {
		// if sd not found... then how about we set it as unknown
		physique.FatCofficient = 0.0
	}

	if found := physique.ResolveBMISD(bmimasters); !found {
		// if bmisd not found...
		physique.BMISD = "Unknown"
	}

	physique.ResolveConclusion()
	return nil
}

// GetPhysiques get physiques
func (c *PhysiqueController) GetPhysiques(class, year, name string) ([]*models.Physique, error) {
	return c.repository.Select(class, year, name)
}

// UpdatePhysique update physique
func (c *PhysiqueController) UpdatePhysique(physique *models.Physique) (err error) {
	pupils, err := c.pupilcontroller.GetPupils(physique.Class, physique.Year)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPupils)
	}

	for _, pupil := range pupils {
		if pupil.Year == physique.Year && pupil.Class == physique.Class && pupil.Name == physique.Name {
			if err = c.ResolvePhysique(physique); err != nil {
				return utils.NewError(errorcode.CoreInvalidPhysique)
			}

			if err = c.repository.Update(physique); err != nil {
				return utils.NewError(errorcode.CoreFailedToUpdatePhysiques)
			}

			return nil
		}
	}

	return utils.NewError(errorcode.CoreInvalidPupil)
}

// UpdatePhysiques delete insert physiques
func (c *PhysiqueController) UpdatePhysiques(physiques []*models.Physique) (err error) {
	foundcount := 0
	pupilsmap := make(map[string][]*models.Pupil)

	for _, physique := range physiques {
		key := fmt.Sprintf("%s_%s", physique.Class, physique.Year)
		pupils, ok := pupilsmap[key]
		if !ok {
			pupils, err = c.pupilcontroller.GetPupils(physique.Class, physique.Year)
			if err != nil {
				return utils.NewError(errorcode.CoreFailedToGetPupils)
			}
			pupilsmap[key] = pupils
		}

		for _, pupil := range pupils {
			if pupil.Year == physique.Year && pupil.Class == physique.Class && pupil.Name == physique.Name {
				foundcount++
			}
		}
	}

	// all the pupils must be in pupils table
	if foundcount != len(physiques) {
		return utils.NewError(errorcode.CoreInvalidPupil)
	}

	for _, physique := range physiques {
		if err = c.ResolvePhysique(physique); err != nil {
			return utils.NewError(errorcode.CoreInvalidPhysique)
		}
	}

	return c.repository.DeleteInsert(physiques)
}

// GetAgeHeightWeightPMasters get age height weight p masters
func (c *PhysiqueController) GetAgeHeightWeightPMasters() ([]*models.AgeHeightWeightPMaster, error) {
	return c.repository.SelectAgeHeightWeightPMasters()
}

// GetAgeHeightWeightSDMasters get age height weight sd masters
func (c *PhysiqueController) GetAgeHeightWeightSDMasters() ([]*models.AgeHeightWeightSDMaster, error) {
	return c.repository.SelectAgeHeightWeightSDMasters()
}

// GetBMIMasters get bmi masters
func (c *PhysiqueController) GetBMIMasters() ([]*models.BMIMaster, error) {
	return c.repository.SelectBMIMasters()
}

// GetHeightToWeightPMasters get height to weight p masters
func (c *PhysiqueController) GetHeightToWeightPMasters() ([]*models.HeightToWeightPMaster, error) {
	return c.repository.SelectHeightToWeightPMasters()
}

// GetHeightToWeightSDMasters get height to weight sd masters
func (c *PhysiqueController) GetHeightToWeightSDMasters() ([]*models.HeightToWeightSDMaster, error) {
	return c.repository.SelectHeightToWeightSDMasters()
}
