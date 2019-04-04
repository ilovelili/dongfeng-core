package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// UserRepository user info repository
type UserRepository struct{}

// NewUserRepository init UserProfile repository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// SelectByMail select user by mail
func (r *UserRepository) SelectByMail(email string) (user *models.User, err error) {
	var u models.User
	query := Table("users").Alias("u").Where().Eq("u.email", email).Eq("u.enabled", 1).Sql()
	if err = session().Find(query, nil).Single(&u); err != nil {
		return
	}
	user = &u

	// then parse the settings
	// step1, get settings master
	var settingsmaster []*models.Settings
	query = Table("settings").Alias("s").Where().Eq("s.enabled", 1).Sql()
	if err = session().Find(query, nil).All(&settingsmaster); err != nil && !norows(err) {
		return
	}

	usersettings := []*models.Settings{}
	for _, settingsitem := range settingsmaster {
		inneritem := &models.Settings{
			ID:      settingsitem.ID,
			Name:    settingsitem.Name,
			Value:   settingsitem.Value,
			Enabled: settingsitem.Value == settingsitem.Value&user.Setting,
		}

		usersettings = append(usersettings, inneritem)
	}
	user.Settings = usersettings

	return
}

// Upsert upsert user and return the ID
func (r *UserRepository) Upsert(user *models.User) error {
	query := countQuery(user.ID)
	count, err := session().Count(query, nil)
	if err != nil || 0 == count {
		return insertTx(user)
	}

	return updateTx(user)
}

// CountByID count by user id
func (r *UserRepository) CountByID(uid string) (int64, error) {
	query := countQuery(uid)
	return session().Count(query, nil)
}

func countQuery(uid string) string {
	return Table("users").Alias("u").Project(`count("u.*")`).Where().Eq("u.id", uid).Sql()
}
