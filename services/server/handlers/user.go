package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// UpdateUser update user info
func (f *Facade) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, rsp *proto.UpdateUserResponse) error {
	if req.GetName() == "" && req.GetAvatar() == "" {
		return utils.NewError(errorcode.CoreInvalidUpdateUserRequest)
	}

	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	if req.GetName() != "" {
		user.Name = req.GetName()
	}

	if req.GetAvatar() != "" {
		user.Avatar = req.GetAvatar()
	}

	usercontroller := controllers.NewUserController()
	err = usercontroller.Save(user)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveUser)
	}

	f.syslog(notification.ProfileUpdated(user.ID))
	return err
}
