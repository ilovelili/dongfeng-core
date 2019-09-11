package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetPhysiques get physiques
func (f *Facade) GetPhysiques(ctx context.Context, req *proto.GetPhysiqueRequest, rsp *proto.GetPhysiqueResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	physiquecontroller := controllers.NewPhysiqueController()
	physiques, err := physiquecontroller.GetPhysiques(req.GetClass(), req.GetYear(), req.GetName())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiques)
	}

	items := []*proto.Physique{}
	for _, physique := range physiques {
		items = append(items, &proto.Physique{
			Id:            physique.ID,
			Year:          physique.Year,
			Class:         physique.Class,
			Name:          physique.Name,
			Gender:        physique.Gender,
			BirthDate:     physique.BirthDate,
			ExamDate:      physique.ExamDate,
			Height:        physique.Height,
			Weight:        physique.Weight,
			Age:           physique.Age,
			AgeCmp:        physique.AgeComparison,
			HeightP:       physique.HeightP,
			WeightP:       physique.WeightP,
			HeightWeightP: physique.HeightToWeightP,
			Bmi:           physique.BMI,
			FatCofficient: physique.FatCofficient,
			Conclusion:    physique.Conclusion,
			CreatedBy:     physique.CreatedBy,
		})
	}

	rsp.Physiques = items
	return nil
}

// UpdatePhysique update physique
func (f *Facade) UpdatePhysique(ctx context.Context, req *proto.UpdatePhysiqueRequest, rsp *proto.UpdatePhysiqueResponse) error {
	pid := req.GetPid()
	userinfo, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	var user *models.User
	err = json.Unmarshal(userinfo, &user)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	physiques := req.GetPhysiques()
	if len(physiques) != 1 {
		return utils.NewError(errorcode.CoreFailedToUpdatePhysiques)
	}

	physique := physiques[0]
	physique.CreatedBy = exsitinguser.Email

	physiquecontroller := controllers.NewPhysiqueController()
	err = physiquecontroller.UpdatePhysique(&models.Physique{
		ID:        physique.GetId(),
		Year:      physique.GetYear(),
		Name:      physique.GetName(),
		Class:     physique.GetClass(),
		Gender:    physique.GetGender(),
		BirthDate: physique.GetBirthDate(),
		ExamDate:  physique.GetExamDate(),
		Height:    physique.GetHeight(),
		Weight:    physique.GetWeight(),
		CreatedBy: physique.GetCreatedBy(),
	})

	f.syslog(notification.PhysiqueUpdated(exsitinguser.ID))
	return err
}

// UpdatePhysiques update physiques
func (f *Facade) UpdatePhysiques(ctx context.Context, req *proto.UpdatePhysiqueRequest, rsp *proto.UpdatePhysiqueResponse) error {
	pid := req.GetPid()
	userinfo, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	var user *models.User
	err = json.Unmarshal(userinfo, &user)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	physiques := []*models.Physique{}
	for _, physique := range req.GetPhysiques() {
		physiques = append(physiques, &models.Physique{
			ID:        physique.GetId(),
			Year:      physique.GetYear(),
			Name:      physique.GetName(),
			Class:     physique.GetClass(),
			Gender:    physique.GetGender(),
			BirthDate: physique.GetBirthDate(),
			ExamDate:  physique.GetExamDate(),
			Height:    physique.GetHeight(),
			Weight:    physique.GetWeight(),
			CreatedBy: exsitinguser.Email,
		})
	}

	physiquecontroller := controllers.NewPhysiqueController()
	err = physiquecontroller.UpdatePhysiques(physiques)

	f.syslog(notification.PhysiqueUpdated(exsitinguser.ID))
	return err
}

// GetAgeHeightWeightPMasters get age height weight p masters
func (f *Facade) GetAgeHeightWeightPMasters(ctx context.Context, req *proto.GetAgeHeightWeightPMasterRequest, rsp *proto.GetAgeHeightWeightPMasterResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	physiquecontroller := controllers.NewPhysiqueController()
	masters, err := physiquecontroller.GetAgeHeightWeightPMasters()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiqueMasters)
	}

	items := []*proto.AgeHeightWeightPMaster{}
	for _, master := range masters {
		items = append(items, &proto.AgeHeightWeightPMaster{
			Id:             master.ID,
			HeightOrWeight: master.HeightOrWeight,
			Gender:         master.Gender,
			AgeMin:         master.AgeMin,
			AgeMax:         master.AgeMax,
			P3:             master.P3,
			P10:            master.P10,
			P20:            master.P20,
			P50:            master.P50,
			P80:            master.P80,
			P97:            master.P97,
		})
	}

	rsp.Masters = items
	return nil
}

// GetAgeHeightWeightSDMasters get age height weight sd masters
func (f *Facade) GetAgeHeightWeightSDMasters(ctx context.Context, req *proto.GetAgeHeightWeightSDMasterRequest, rsp *proto.GetAgeHeightWeightSDMasterResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	physiquecontroller := controllers.NewPhysiqueController()
	masters, err := physiquecontroller.GetAgeHeightWeightSDMasters()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiqueMasters)
	}

	items := []*proto.AgeHeightWeightSDMaster{}
	for _, master := range masters {
		items = append(items, &proto.AgeHeightWeightSDMaster{
			Id:             master.ID,
			HeightOrWeight: master.HeightOrWeight,
			Gender:         master.Gender,
			Age:            master.Age,
			Sdm2:           master.SDM2,
			Sdm1:           master.SDM1,
			Avg:            master.Average,
			Sd1:            master.SD1,
			Sd2:            master.SD2,
		})
	}

	rsp.Masters = items
	return nil
}

// GetBMIMasters get bmi masters
func (f *Facade) GetBMIMasters(ctx context.Context, req *proto.GetBMIMasterRequest, rsp *proto.GetBMIMasterResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	physiquecontroller := controllers.NewPhysiqueController()
	masters, err := physiquecontroller.GetBMIMasters()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiqueMasters)
	}

	items := []*proto.BMIMaster{}
	for _, master := range masters {
		items = append(items, &proto.BMIMaster{
			Id:     master.ID,
			Gender: master.Gender,
			Age:    master.Age,
			Avg:    master.Average,
			Sd1:    master.SD1,
			Sd2:    master.SD2,
			Sd3:    master.SD3,
		})
	}

	rsp.Masters = items
	return nil
}

// GetHeightToWeightPMasters get height to weight p masters
func (f *Facade) GetHeightToWeightPMasters(ctx context.Context, req *proto.GetHeightToWeightPMasterRequest, rsp *proto.GetHeightToWeightPMasterResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	physiquecontroller := controllers.NewPhysiqueController()
	masters, err := physiquecontroller.GetHeightToWeightPMasters()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiqueMasters)
	}

	items := []*proto.HeightToWeightPMaster{}
	for _, master := range masters {
		items = append(items, &proto.HeightToWeightPMaster{
			Id:     master.ID,
			Gender: master.Gender,
			Height: master.Height,
			P3:     master.P3,
			P10:    master.P10,
			P20:    master.P20,
			P50:    master.P50,
			P80:    master.P80,
			P97:    master.P97,
		})
	}

	rsp.Masters = items
	return nil
}

// GetHeightToWeightSDMasters get height to weight sd masters
func (f *Facade) GetHeightToWeightSDMasters(ctx context.Context, req *proto.GetHeightToWeightSDMasterRequest, rsp *proto.GetHeightToWeightSDMasterResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	physiquecontroller := controllers.NewPhysiqueController()
	masters, err := physiquecontroller.GetHeightToWeightSDMasters()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPhysiqueMasters)
	}

	items := []*proto.HeightToWeightSDMaster{}
	for _, master := range masters {
		items = append(items, &proto.HeightToWeightSDMaster{
			Id:     master.ID,
			Gender: master.Gender,
			Height: master.Height,
			Sdm3:   master.SDM3,
			Sdm2:   master.SDM2,
			Sdm1:   master.SDM1,
			Sd0:    master.SD0,
			Sd1:    master.SD1,
			Sd2:    master.SD2,
			Sd3:    master.SD3,
		})
	}

	rsp.Masters = items
	return nil
}
