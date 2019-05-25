package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// ProfileRepository profile repository
type ProfileRepository struct{}

// NewProfileRepository init profile repository
func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

// Select select profiles
func (r *ProfileRepository) Select(year, class, name string) (profile *models.Profile, err error) {
	query := Table("profiles").Alias("p").Where().
		Eq("p.year", year).
		Eq("p.class", class).
		Eq("p.name", name).
		Sql()

	var p models.Profile
	if err = session().Find(query, nil).Single(&p); err != nil && norows(err) {
		err = nil
	}

	if err != nil {
		return
	}

	profile = &p
	return
}

// Upsert upsert profile
func (r *ProfileRepository) Upsert(profile *models.Profile) (err error) {
	query := Table("profiles").Alias("p").Project("p.id").Where().
		Eq("p.year", profile.Year).
		Eq("p.class", profile.Class).
		Eq("p.name", profile.Name).
		Sql()

	var id int64
	err = session().Find(query, nil).Scalar(&id)
	if err != nil || 0 == id {
		err = session().Insert(profile)
	} else {
		profile.ID = id
		err = session().Update(profile)
	}

	return
}
