package repositories

import (
	"fmt"
	"strings"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// PhysiqueRepository physique repository
type PhysiqueRepository struct{}

// NewPhysiqueRepository init physique repository
func NewPhysiqueRepository() *PhysiqueRepository {
	return &PhysiqueRepository{}
}

// SelectAgeHeightWeightPMasters select age weight / height p master
func (r *PhysiqueRepository) SelectAgeHeightWeightPMasters() (pmasters []*models.AgeHeightWeightPMaster, err error) {
	query := Table("physique_height_to_weight_p_master").Sql()
	err = session().Find(query, nil).All(&pmasters)
	return
}

// SelectAgeHeightWeightSDMasters select age weight / height sd master
func (r *PhysiqueRepository) SelectAgeHeightWeightSDMasters() (sdmasters []*models.AgeHeightWeightSDMaster, err error) {
	query := Table("physique_age_height_weight_sd_master").Sql()
	err = session().Find(query, nil).All(&sdmasters)
	return
}

// SelectHeightToWeightPMasters select height weight p masters
func (r *PhysiqueRepository) SelectHeightToWeightPMasters() (hwpmasters []*models.HeightToWeightPMaster, err error) {
	query := Table("physique_height_to_weight_p_master").Sql()
	err = session().Find(query, nil).All(&hwpmasters)
	return
}

// SelectHeightToWeightSDMasters select height to weight sd masters
func (r *PhysiqueRepository) SelectHeightToWeightSDMasters() (hwsdmasters []*models.HeightToWeightSDMaster, err error) {
	query := Table("physique_height_to_weight_sd_master").Sql()
	err = session().Find(query, nil).All(&hwsdmasters)
	return
}

// SelectBMIMasters select bmi sd master
func (r *PhysiqueRepository) SelectBMIMasters() (bmimasters []*models.BMIMaster, err error) {
	query := Table("physique_bmi_master").Sql()
	err = session().Find(query, nil).All(&bmimasters)
	return
}

// Select physiques
func (r *PhysiqueRepository) Select(year, class, name string) (physiques []*models.Physique, err error) {
	querybuilder := Table("physiques").Alias("p").Where()
	var query string

	if class == "" && year == "" && name == "" {
		querybuilder = querybuilder.Eq("1", "1")
	} else {
		if class != "" {
			querybuilder = querybuilder.Eq("p.class", class)
		}

		if year != "" {
			querybuilder = querybuilder.Eq("p.year", year)
		}

		if name != "" {
			querybuilder = querybuilder.Eq("p.name", name)
		}
	}

	query = querybuilder.Sql()
	err = session().Find(query, nil).All(&physiques)
	return
}

// Update update physique
func (r *PhysiqueRepository) Update(physique *models.Physique) (err error) {
	return session().Update(physique)
}

// DeleteInsert delete insert physiques
func (r *PhysiqueRepository) DeleteInsert(physiques []*models.Physique) (err error) {
	physiquesmap := make(map[string][]*models.Physique)
	for _, physique := range physiques {
		key := fmt.Sprintf("%s_%s", physique.Year, physique.Class)
		if v, ok := physiquesmap[key]; !ok {
			physiquesmap[key] = []*models.Physique{physique}
		} else {
			physiquesmap[key] = append(v, physique)
		}
	}

	tx, err := session().Begin()
	if err != nil {
		return
	}

	for k := range physiquesmap {
		segments := strings.Split(k, "_")
		if len(segments) != 2 {
			err = fmt.Errorf("invalid key")
			return
		}

		year, class := segments[0], segments[1]
		_, err = session().ExecTx(tx, fmt.Sprintf("CALL spDeletePhysiques('%s','%s')", year, class))
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	for _, physique := range physiques {
		err = session().InsertTx(tx, physique)
		if err != nil {
			session().Rollback(tx)
			return
		}
	}

	return session().Commit(tx)
}
