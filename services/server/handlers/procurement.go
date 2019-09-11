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

// GetProcurements get procurement
func (f *Facade) GetProcurements(ctx context.Context, req *proto.GetProcurementRequest, rsp *proto.GetProcurementResponse) error {
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

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	procurementcontroller := controllers.NewProcurementController()
	from, to := req.GetFrom(), req.GetTo()

	procurements, err := procurementcontroller.GetProcurements(from, to)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetProcurement)
	}

	rsp.Procurements = procurements
	return nil
}

// UpdateProcurement update procurement
func (f *Facade) UpdateProcurement(ctx context.Context, req *proto.UpdateProcurementRequest, rsp *proto.UpdateProcurementResponse) error {
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

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	procurementcontroller := controllers.NewProcurementController()
	id, amount := req.GetId(), req.GetAmount()
	err = procurementcontroller.UpdateRecipeUnitAmount(&models.Recipe{
		ID:         id,
		UnitAmount: amount,
	})

	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateProcurement)
	}

	return err
}
