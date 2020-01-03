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
func (r *RoleRepository) Select(user string) (access []string, err error) {
	query := Table("roles").Alias("r").Where().Eq("r.user", user).Sql()
	var roles []*models.Role
	if err := session().Find(query, nil).All(&roles); norows(err) {
		err = nil
		access = []string{"班级信息", "园儿信息"}
	} else {
		for _, r := range roles {
			access = append(access, r.Path)
		}
	}

	return
}
