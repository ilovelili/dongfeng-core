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

// GetClasses get classes
func (f *Facade) GetClasses(ctx context.Context, req *proto.GetClassRequest, rsp *proto.GetClassResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
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

	classes := req.GetClasses()
	for _, class := range classes {
		class.CreatedBy = exsitinguser.Email
	}

	classcontroller := controllers.NewClassController()
	err = classcontroller.UpdateClasses(classes)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateClasses)
	}

	f.syslog(notification.ClasslistUpdated(exsitinguser.ID))
	return nil
}
