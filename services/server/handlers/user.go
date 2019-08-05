package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"	
)

// UpdateUser update user info
func (f *Facade) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, rsp *proto.UpdateUserResponse) error {
	if req.GetName() == "" && req.GetAvatar() == "" {
		return utils.NewError(errorcode.CoreInvalidUpdateUserRequest)
	}

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

	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	if req.GetName() != "" {
		user.Name = req.GetName()
	}

	if req.GetAvatar() != "" {
		user.Avatar = req.GetAvatar()
	}

	err = usercontroller.Save(user)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveUser)
	}

	f.syslog(notification.ProfileUpdated(user.ID))
	return err
}
