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

// GetProfileTemplate get profile tempalte
func (f *Facade) GetProfileTemplate(ctx context.Context, req *proto.GetProfileTemplateRequest, rsp *proto.GetProfileTemplateResponse) error {
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
	template, err := profiletemplatecontroller.GetProfileTemplate(req.GetName())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetGrowthProfile)
	}

	rsp.Profile = template.Profile
	return nil
}

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
		return utils.NewError(errorcode.CoreFailedToGetGrowthProfile)
	}

	_templates := []*proto.ProfileTemplate{}
	for _, template := range templates {
		_templates = append(_templates, &proto.ProfileTemplate{
			Name:      template.Name,
			CreatedBy: template.CreatedBy,
		})
	}

	rsp.Templates = _templates
	return nil
}

// UpdateProfileTemplate update profile tempalte
func (f *Facade) UpdateProfileTemplate(ctx context.Context, req *proto.UpdateProfileTemplateRequest, rsp *proto.UpdateProfileTemplateResponse) error {
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
	err = profiletemplatecontroller.UpdateProfileTemplates(&models.ProfileTemplate{
		Name:      req.GetName(),
		Enabled:   req.GetEnabled(),
		Profile:   req.GetProfile(),
		CreatedBy: user.Email,
	})
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveProfileTemplate)
	}

	return nil
}
