package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// ProfileTemplateRepository profile template repository
type ProfileTemplateRepository struct{}

// NewProfileTemplateRepository init profile template repository
func NewProfileTemplateRepository() *ProfileTemplateRepository {
	return &ProfileTemplateRepository{}
}

// Select select ProfileTemplates
func (r *ProfileTemplateRepository) Select() (templates []*models.ProfileTemplate, err error) {
	query := Table("profile_templates").Alias("g").Sql()
	// no rows is actually not an error
	if err = session().Find(query, nil).All(&templates); err != nil && norows(err) {
		err = nil
	}
	return
}

// Upsert upsert  profile template
func (r *ProfileTemplateRepository) Upsert(template *models.ProfileTemplate) (err error) {
	query := Table("profile_templates").Alias("g").Project("g.id").Sql()
	var id int64
	err = session().Find(query, nil).Scalar(&id)
	if err != nil || 0 == id {
		err = session().Insert(template)
	} else {
		template.ID = id
		err = session().Update(template)
	}
	return
}
