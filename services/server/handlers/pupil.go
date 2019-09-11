package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// GetPupils get pupils
func (f *Facade) GetPupils(ctx context.Context, req *proto.GetPupilRequest, rsp *proto.GetPupilResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
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
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	pupils := req.GetPupils()
	if len(pupils) != 1 {
		return utils.NewError(errorcode.CoreFailedToUpdatePupils)
	}

	pupil := pupils[0]
	pupil.CreatedBy = exsitinguser.Email

	pupilcontroller := controllers.NewPupilController()
	err = pupilcontroller.UpdatePupil(&models.Pupil{
		ID:        pupil.GetId(),
		Year:      pupil.GetYear(),
		Name:      pupil.GetName(),
		Class:     pupil.GetClass(),
		CreatedBy: pupil.CreatedBy,
	})

	f.syslog(notification.NamelistUpdated(exsitinguser.ID))
	return err
}

// UpdatePupils update pupils
func (f *Facade) UpdatePupils(ctx context.Context, req *proto.UpdatePupilRequest, rsp *proto.UpdatePupilResponse) error {
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
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	pupils := req.GetPupils()
	for _, pupil := range pupils {
		pupil.CreatedBy = exsitinguser.Email
	}

	pupilcontroller := controllers.NewPupilController()
	err = pupilcontroller.UpdatePupils(pupils)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdatePupils)
	}

	f.syslog(notification.NamelistUpdated(exsitinguser.ID))
	return nil
}
