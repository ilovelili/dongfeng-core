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

// GetProfileTemplates get profile tempaltes
func (f *Facade) GetProfileTemplates(ctx context.Context, req *proto.GetProfileTemplatesRequest, rsp *proto.GetProfileTemplatesResponse) error {
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

	profiletemplatecontroller := controllers.NewProfileTemplateController()
	templates, err := profiletemplatecontroller.GetProfileTemplates()
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToConvertEbookHTML)
	}

	_templates := []*proto.ProfileTemplate{}
	for _, template := range templates {
		_templates = append(_templates, &proto.ProfileTemplate{
			Id:        template.ID,
			Name:      template.Name,
			CreatedBy: template.CreatedBy,
		})
	}

	rsp.Templates = _templates
	return nil
}
