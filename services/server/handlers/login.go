// Package handlers define the core behaviors of each API
package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// Login handler returns all data needed by front end
func (f *Facade) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
	pid := req.GetPid()
	userinfo, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	var user *models.User
	err = json.Unmarshal(userinfo, &user)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// user email empty caused by open ID login
	if user.Email == "" {
		user.Email = fmt.Sprintf("%s@dongfeng.cn", pid)
	}

	// check if user exists or not
	var newUser bool
	usercontroller := controllers.NewUserController()
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	newUser = err != nil

	if newUser {
		// new user, save to database
		if err = usercontroller.Save(user); err != nil {
			return utils.NewError(errorcode.CoreFailedToSaveUser)
		}
	} else {
		user = exsitinguser
	}

	// user profile
	rsp.User = &proto.User{
		Newuser:  newUser,
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Settings: resolveSettings(user.Settings),
	}

	return err
}

func resolveSettings(settings []*models.Settings) []*proto.Setting {
	result := []*proto.Setting{}
	for _, setting := range settings {
		result = append(result, &proto.Setting{
			Id:      setting.ID,
			Name:    setting.Name,
			Enabled: setting.Enabled,
		})
	}

	return result
}
