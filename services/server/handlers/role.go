package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetRole get role
func (f *Facade) GetRole(ctx context.Context, req *proto.GetRoleRequest, rsp *proto.GetRoleResponse) error {
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

	rolecontroller := controllers.NewRoleController()
	role, err := rolecontroller.GetRole(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetRole)
	}

	rsp.Role = role
	return nil
}
