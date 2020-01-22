package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetRole get role
func (f *Facade) GetRole(ctx context.Context, req *proto.GetRoleRequest, rsp *proto.GetRoleResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	rolecontroller := controllers.NewRoleController()
	role, err := rolecontroller.GetRole(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetRole)
	}

	rsp.Role = role
	return nil
}
