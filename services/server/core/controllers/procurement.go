package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// ProcurementController notification controller
type ProcurementController struct {
	attendancecontroller *AttendanceController
	menucontroller       *MenuController
	recipecontroller     *RecipeController
}

// NewProcurementController new controller
func NewProcurementController() *ProcurementController {
	return &ProcurementController{
		attendancecontroller: NewAttendanceController(),
		menucontroller:       NewMenuController(),
		recipecontroller:     NewRecipeController(),
	}
}

// GetProcurements get procurements
func (c *ProcurementController) GetProcurements(from, to string) (procurements []*proto.Procurement, err error) {
	// date recipe name map
	datarecipenamemap := make(map[string][]string)
	// date / ingredient nutrition map
	datarecipemap := make(map[string][]*models.RecipeCombined)
	// date / attendance map
	dateattendancemap := make(map[string]int64)

	menus, menuerr := c.menucontroller.GetMenus(from, to, -1, -1)
	if menuerr != nil {
		err = menuerr
		return
	}

	for _, menu := range menus {
		if v, ok := datarecipenamemap[menu.Date]; !ok {
			datarecipenamemap[menu.Date] = []string{menu.Recipe}
		} else {
			contains := false
			for _, r := range v {
				if r == menu.Recipe {
					contains = true
					break
				}
			}

			if !contains {
				datarecipenamemap[menu.Date] = append(v, menu.Recipe)
			}
		}
	}

	for date, recipenames := range datarecipenamemap {
		recipes, reciperr := c.recipecontroller.SelectRecipes(recipenames)
		if reciperr != nil {
			err = reciperr
			return
		}

		for _, recipe := range recipes {
			if v, ok := datarecipemap[date]; !ok {
				datarecipemap[date] = []*models.RecipeCombined{recipe}
			} else {
				contains := false
				for _, r := range v {
					if r.Name == recipe.Name && r.IngredientName == recipe.IngredientName {
						contains = true
						break
					}
				}

				if !contains {
					datarecipemap[date] = append(v, recipe)
				}
			}
		}
	}

	attendances, attendancerr := c.attendancecontroller.CountAttendance(from, to, "")
	if attendancerr != nil {
		err = attendancerr
		return
	}

	// tbd. check attendance count implementation
	for _, attendance := range attendances.ClassAttendanceCounts {
		if v, ok := dateattendancemap[attendance.Date]; !ok {
			dateattendancemap[attendance.Date] = attendance.Count
		} else {
			dateattendancemap[attendance.Date] = v + attendance.Count
		}
	}

	procurements = make([]*proto.Procurement, 0)
	for date := range datarecipenamemap {
		procurement := &proto.Procurement{
			Date: date,
		}

		if v, ok := dateattendancemap[date]; !ok {
			// this should not happen
			continue
		} else {
			procurement.Attendance = v
		}

		if v, ok := datarecipemap[date]; !ok {
			// this should not happen
			continue
		} else {
			ingredientunitamounts := make([]*proto.IngredientAmount, 0)
			for _, r := range v {
				ingredientunitamounts = append(ingredientunitamounts, &proto.IngredientAmount{
					Ingredient: r.IngredientName,
					Amount:     r.UnitAmount * float64(procurement.Attendance),
					Matched:    r.Category != "",
				})
			}
			procurement.Procurements = ingredientunitamounts
		}

		procurements = append(procurements, procurement)
	}

	return
}
