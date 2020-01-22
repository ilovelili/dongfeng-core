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

// GetPupils get pupils
func (f *Facade) GetPupils(ctx context.Context, req *proto.GetPupilRequest, rsp *proto.GetPupilResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	pupilcontroller := controllers.NewPupilController()
	pupils, err := pupilcontroller.GetPupils(req.GetClass(), req.GetYear())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetPupils)
	}

	items := []*proto.Pupil{}
	for _, pupil := range pupils {
		items = append(items, &proto.Pupil{
			Id:        pupil.ID,
			Year:      pupil.Year,
			Class:     pupil.Class,
			Name:      pupil.Name,
			CreatedBy: pupil.CreatedBy,
		})
	}

	rsp.Pupils = items
	return nil
}

// UpdatePupil update pupil
func (f *Facade) UpdatePupil(ctx context.Context, req *proto.UpdatePupilRequest, rsp *proto.UpdatePupilResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	pupils := req.GetPupils()
	if len(pupils) != 1 {
		return utils.NewError(errorcode.CoreFailedToUpdatePupils)
	}

	pupil := pupils[0]
	pupil.CreatedBy = user.Email

	pupilcontroller := controllers.NewPupilController()
	err = pupilcontroller.UpdatePupil(&models.Pupil{
		ID:        pupil.GetId(),
		Year:      pupil.GetYear(),
		Name:      pupil.GetName(),
		Class:     pupil.GetClass(),
		CreatedBy: pupil.CreatedBy,
	})

	f.syslog(notification.NamelistUpdated(user.ID))
	return err
}

// UpdatePupils update pupils
func (f *Facade) UpdatePupils(ctx context.Context, req *proto.UpdatePupilRequest, rsp *proto.UpdatePupilResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	pupils := req.GetPupils()
	for _, pupil := range pupils {
		pupil.CreatedBy = user.Email
	}

	pupilcontroller := controllers.NewPupilController()
	err = pupilcontroller.UpdatePupils(pupils)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdatePupils)
	}

	f.syslog(notification.NamelistUpdated(user.ID))
	return nil
}
