package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// RoleRepository role repository
type RoleRepository struct{}

// NewRoleRepository init role repository
func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

// Select select access paths by user
func (r *RoleRepository) Select(user string) (role *models.Role, err error) {
	query := Table("roles").Alias("r").Where().Eq("r.user", user).Sql()

	var _role models.Role
	if err := session().Find(query, nil).Single(&_role); norows(err) {
		err = nil
		role = &models.Role{
			User: user,
			Role: "default",
		}
		return role, err
	}

	role = &_role
	return
}
