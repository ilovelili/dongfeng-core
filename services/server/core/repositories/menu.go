package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// MenuRepository menu repository
type MenuRepository struct{}

// NewMenuRepository init menu repository
func NewMenuRepository() *MenuRepository {
	return &MenuRepository{}
}

// Select select Menus
func (r *MenuRepository) Select(from string, to string, mealid int64, juniororsenior int64) (menus []*models.Menu, err error) {
	querybuilder := Table("menus").Alias("m").Where().
		Gte("m.date", from).
		Lte("m.date", to).
		Ne("m.recipe", "未排菜") // exclude "未排菜"

	if mealid == 0 || mealid == 1 || mealid == 2 {
		querybuilder = querybuilder.Eq("m.breakfast_or_lunch", mealid)
	}

	if juniororsenior == 0 || juniororsenior == 1 {
		querybuilder = querybuilder.Eq("m.junior_or_senior_class", juniororsenior)
	}

	query := querybuilder.Sql()
	if err = session().Find(query, nil).All(&menus); err != nil && norows(err) {
		err = nil
	}

	return
}
