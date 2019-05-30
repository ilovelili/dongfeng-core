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
func (r *ProfileRepository) Select(year, class, name, date string) (profile *models.Profile, err error) {
	query := Table("profiles").Alias("p").Where().
		Eq("p.year", year).
		Eq("p.class", class).
		Eq("p.name", name).
		Eq("p.date", date).
		Eq("p.enabled", 1).
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

// SelectAll select all profiles
func (r *ProfileRepository) SelectAll(year, class, name string) (profiles []*models.Profile, err error) {
	querybuilder := Table("profiles").Alias("p").Where().Eq("p.enabled", 1)

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

	query := querybuilder.Project(
		"p.id",
		"p.year",
		"p.class",
		"p.date",
		"p.name",
	).Sql()

	if err = session().Find(query, nil).All(&profiles); err != nil && norows(err) {
		err = nil
	}

	return
}

// Upsert upsert profile
func (r *ProfileRepository) Upsert(profile *models.Profile) (err error) {
	query := Table("profiles").Alias("p").Project("p.id").Where().
		Eq("p.year", profile.Year).
		Eq("p.class", profile.Class).
		Eq("p.name", profile.Name).
		Eq("p.date", profile.Date).
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
