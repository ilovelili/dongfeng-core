package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetMenus get menus
func (f *Facade) GetMenus(ctx context.Context, req *proto.GetMenuRequest, rsp *proto.GetMenuResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	claims, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// Unmarshal user info
	userinfo, _ := json.Marshal(claims)
	var user *models.User
	err = json.Unmarshal(userinfo, &user)

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	menucontroller := controllers.NewMenuController()
	menus, err := menucontroller.GetMenus(req.GetFrom(), req.GetTo(), req.GetBreakfastOrLunch(), req.GetJuniorOrSenior())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetMenu)
	}

	menumap := make(map[string][]string)
	for _, menu := range menus {
		key := fmt.Sprintf("%s_%d_%d", menu.Date, menu.BreakfastOrLunch, menu.JuniorOrSenior)
		if recipes, ok := menumap[key]; !ok {
			menumap[key] = []string{menu.Recipe}
		} else {
			menumap[key] = append(recipes, menu.Recipe)
		}
	}

	m := []*proto.Menu{}
	for k, v := range menumap {
		segments := strings.Split(k, "_")
		if len(segments) != 3 {
			return utils.NewError(errorcode.CoreFailedToGetMenu)
		}

		date, breakfast_or_lunch_str, junior_or_senior_str := segments[0], segments[1], segments[2]
		breakfast_or_lunch, _ := strconv.ParseInt(breakfast_or_lunch_str, 10, 64)
		junior_or_senior, _ := strconv.ParseInt(junior_or_senior_str, 10, 64)

		m = append(m, &proto.Menu{
			Date:             date,
			BreakfastOrLunch: breakfast_or_lunch,
			JuniorOrSenior:   junior_or_senior,
			Recipe:           resolveRecipes(v),
		})
	}

	rsp.Menus = m
	return nil
}

// resolveRecipes join all the recipes by comma and remove duplications in case
func resolveRecipes(recipes []string) string {
	recipes = sharedlib.Unique(recipes)
	return strings.Join(recipes, ",")
}
