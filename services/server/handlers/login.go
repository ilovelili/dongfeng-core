// Package handlers define the core behaviors of each API
package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// Login handler returns all data needed by front end
func (f *Facade) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
	pid, email, newUser := req.GetPid(), req.GetEmail(), false
	user, err := f.parseUser(pid, email)
	if err != nil {
		if err.Error() == utils.NewError(errorcode.CoreNoUser).Error() {
			newUser = true
		} else {
			return err
		}
	}

	usercontroller := controllers.NewUserController()
	if newUser {
		// new user, save to database
		if err = usercontroller.Save(user); err != nil {
			return utils.NewError(errorcode.CoreFailedToSaveUser)
		}
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
