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

// SelectAll select all templates
func (r *ProfileTemplateRepository) SelectAll() (templates []*models.ProfileTemplate, err error) {
	query := Table("profile_templates").Alias("g").Where().Eq("g.enabled", 1).Sql()
	// no rows is actually not an error
	if err = session().Find(query, nil).All(&templates); err != nil && norows(err) {
		err = nil
	}
	return
}

// Select select template by name
func (r *ProfileTemplateRepository) Select(name string) (template *models.ProfileTemplate, err error) {
	query := Table("profile_templates").Alias("g").Where().Eq("g.name", name).Eq("g.enabled", 1).Sql()
	// no rows is actually not an error
	var _template models.ProfileTemplate
	if err = session().Find(query, nil).Single(&_template); err != nil && norows(err) {
		err = nil
		return
	}

	template = &_template
	return
}

// Upsert upsert  profile template
func (r *ProfileTemplateRepository) Upsert(template *models.ProfileTemplate) (err error) {
	query := Table("profile_templates").Alias("g").Project("g.id").Where().Eq("g.name", template.Name).Sql()
	var id int64
	err = session().Find(query, nil).Scalar(&id)
	if err != nil || 0 == id {
		err = session().Insert(template)
	} else {
		err = session().Update(template)
	}

	return
}
