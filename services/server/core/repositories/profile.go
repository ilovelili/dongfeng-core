package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// ProfileRepository profile repository
type ProfileRepository struct{}

// NewProfileRepository init profile repository
func NewProfileRepository() *ProfileRepository {
	return &ProfileRepository{}
}

// Select select profile
func (r *ProfileRepository) Select(year, class, name, date string) (profile *models.Profile, err error) {
	query := Table("profiles").Alias("p").Where().
		Eq("p.year", year).
		Eq("p.class", class).
		Eq("p.name", name).
		Eq("p.date", date).
		Eq("p.enabled", 1).
		Sql()

	var p models.Profile
	if err = session().Find(query, nil).Single(&p); err != nil && norows(err) || p.Profile == "" {
		query = Table("profiles").Alias("p").Where().
			Eq("p.year", year).
			Eq("p.class", class).
			Eq("p.name", class).
			Eq("p.date", date).
			Eq("p.enabled", 1).
			Sql()

		if err = session().Find(query, nil).Single(&p); err != nil && norows(err) {
			err = nil
		}
	}

	if err != nil {
		return
	}

	profile = &p
	return
}

// SelectPrev select previous profile
func (r *ProfileRepository) SelectPrev(year, class, name, date string) (profile *models.Profile, err error) {
	query := Table("profiles").Alias("p").Where().
		Eq("p.year", year).
		Eq("p.class", class).
		Eq("p.name", name).
		Lt("p.date", date).
		Eq("p.enabled", 1).
		Take(1).
		Sql()

	var p models.Profile
	if err = session().Find(query, nil).Single(&p); err != nil && norows(err) || p.Profile == "" {
		query = Table("profiles").Alias("p").Where().
			Eq("p.year", year).
			Eq("p.class", class).
			Eq("p.name", class).
			Lt("p.date", date).
			Eq("p.enabled", 1).
			Sql()

		if err = session().Find(query, nil).Single(&p); err != nil && norows(err) {
			err = nil
		}
	}

	if err != nil {
		return
	}

	profile = &p
	return
}

// SelectNext select next profile
func (r *ProfileRepository) SelectNext(year, class, name, date string) (profile *models.Profile, err error) {
	query := Table("profiles").Alias("p").Where().
		Eq("p.year", year).
		Eq("p.class", class).
		Eq("p.name", name).
		Gt("p.date", date).
		Eq("p.enabled", 1).
		Take(1).
		Sql()

	var p models.Profile
	if err = session().Find(query, nil).Single(&p); err != nil && norows(err) || p.Profile == "" {
		query = Table("profiles").Alias("p").Where().
			Eq("p.year", year).
			Eq("p.class", class).
			Eq("p.name", class).
			Gt("p.date", date).
			Eq("p.enabled", 1).
			Sql()

		if err = session().Find(query, nil).Single(&p); err != nil && norows(err) {
			err = nil
		}
	}

	if err != nil {
		return
	}

	profile = &p
	return
}

// SelectAll select all profiles
func (r *ProfileRepository) SelectAll(year, class, name string) (profiles []*models.Profile, err error) {
	querybuilder := Table("profiles").Alias("p").Where().Ne("p.name", class).Eq("p.enabled", 1)
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

// Insert insert profile
func (r *ProfileRepository) Insert(profile *models.Profile) (err error) {
	var profiles []*models.Profile
	query := Table("profiles").Alias("p").Where().
		Eq("p.year", profile.Year).
		Eq("p.class", profile.Class).
		Eq("p.name", profile.Name).
		Eq("p.date", profile.Date).
		Sql()

	if err = session().Find(query, nil).All(&profiles); err != nil {
		if norows(err) {
			// not exist, so insert
			profile.Enabled = true
			return session().Insert(profile)
		}
		return
	} else if len(profiles) == 0 {
		profile.Enabled = true
		return session().Insert(profile)
	}

	for _, p := range profiles {
		// already exists a valid profile, return error
		if p.Enabled {
			return fmt.Errorf("already exist")
		}
	}

	// get the first disabled profile (there should be only one) and update
	_profile := profiles[0]
	profile.ID = _profile.ID
	profile.Profile = "{}"
	profile.Enabled = true
	return session().Update(profile)
}

// Delete delete profile by set enabled to false
func (r *ProfileRepository) Delete(profile *models.Profile) (err error) {
	query := Table("profiles").Alias("p").Where().
		Eq("p.year", profile.Year).
		Eq("p.class", profile.Class).
		Eq("p.name", profile.Name).
		Eq("p.date", profile.Date).
		Sql()

	var p models.Profile
	err = session().Find(query, nil).Single(&p)
	// not found which is good
	if err != nil || 0 == p.ID {
		return nil
	}

	profile.ID = p.ID
	profile.Profile = p.Profile
	profile.Enabled = false
	return session().Update(profile)
}
