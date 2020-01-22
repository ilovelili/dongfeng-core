package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetClasses get classes
func (f *Facade) GetClasses(ctx context.Context, req *proto.GetClassRequest, rsp *proto.GetClassResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	classcontroller := controllers.NewClassController()
	classes, err := classcontroller.GetClasses()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetClasses)
	}

	items := []*proto.Class{}
	for _, class := range classes {
		items = append(items, &proto.Class{
			Id:        class.ID,
			Name:      class.Name,
			CreatedBy: class.CreatedBy,
		})
	}

	rsp.Classes = items
	return nil
}

// UpdateClasses update classes
func (f *Facade) UpdateClasses(ctx context.Context, req *proto.UpdateClassRequest, rsp *proto.UpdateClassResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	classes := req.GetClasses()
	for _, class := range classes {
		class.CreatedBy = user.Email
	}

	classcontroller := controllers.NewClassController()
	err = classcontroller.UpdateClasses(classes)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateClasses)
	}

	f.syslog(notification.ClasslistUpdated(user.ID))
	return nil
}
