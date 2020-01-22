package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetProcurements get procurement
func (f *Facade) GetProcurements(ctx context.Context, req *proto.GetProcurementRequest, rsp *proto.GetProcurementResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
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
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
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

	f.syslog(notification.ProcurementUpdated(user.ID))
	return err
}
