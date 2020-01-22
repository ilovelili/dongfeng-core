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

// GetProfile get profile
func (f *Facade) GetProfile(ctx context.Context, req *proto.GetProfileRequest, rsp *proto.GetProfileResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	profilecontroller := controllers.NewProfileController()
	profile, err := profilecontroller.GetProfile(req.GetYear(), req.GetClass(), req.GetName(), req.GetDate())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetGrowthProfile)
	}

	rsp.Profile = profile.Profile
	return nil
}

// GetPrevProfile get previous profile
func (f *Facade) GetPrevProfile(ctx context.Context, req *proto.GetPrevOrNextProfileRequest, rsp *proto.GetPrevOrNextProfileResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	profilecontroller := controllers.NewProfileController()
	profile, err := profilecontroller.GetPrevProfile(req.GetYear(), req.GetClass(), req.GetName(), req.GetDate())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetGrowthProfile)
	}

	rsp.Date = profile.Date
	return nil
}

// GetNextProfile get next profile
func (f *Facade) GetNextProfile(ctx context.Context, req *proto.GetPrevOrNextProfileRequest, rsp *proto.GetPrevOrNextProfileResponse) error {
	pid := req.GetPid()
	_, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	profilecontroller := controllers.NewProfileController()
	profile, err := profilecontroller.GetNextProfile(req.GetYear(), req.GetClass(), req.GetName(), req.GetDate())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetGrowthProfile)
	}

	rsp.Date = profile.Date
	return nil
}

// GetProfiles get profile
func (f *Facade) GetProfiles(ctx context.Context, req *proto.GetProfilesRequest, rsp *proto.GetProfilesResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	profilecontroller := controllers.NewProfileController()
	profiles, err := profilecontroller.GetProfiles(req.GetYear(), req.GetClass(), req.GetName())
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetGrowthProfile)
	}

	_profiles := []*proto.Profile{}
	for _, profile := range profiles {
		_profiles = append(_profiles, &proto.Profile{
			Id:    profile.ID,
			Year:  profile.Year,
			Class: profile.Class,
			Name:  profile.Name,
			Date:  profile.Date,
		})
	}

	rsp.Profiles = _profiles
	return nil
}

// UpdateProfile update profile
func (f *Facade) UpdateProfile(ctx context.Context, req *proto.UpdateProfileRequest, rsp *proto.UpdateProfileResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	profilecontroller := controllers.NewProfileController()
	err = profilecontroller.SaveProfile(&models.Profile{
		Year:      req.GetYear(),
		Class:     req.GetClass(),
		Name:      req.GetName(),
		Date:      req.GetDate(),
		Profile:   req.GetProfile(),
		Enabled:   req.GetEnabled(),
		CreatedBy: user.Email,
	})

	f.syslog(notification.GrowthProfileUpdated(user.ID))
	return err
}

// CreateProfile create profile
func (f *Facade) CreateProfile(ctx context.Context, req *proto.UpdateProfileRequest, rsp *proto.UpdateProfileResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	profilecontroller := controllers.NewProfileController()
	err = profilecontroller.InsertProfile(&models.Profile{
		Year:      req.GetYear(),
		Class:     req.GetClass(),
		Name:      req.GetName(),
		Date:      req.GetDate(),
		Template:  req.GetTemplateName(),
		CreatedBy: user.Email,
	})

	f.syslog(notification.GrowthProfileUpdated(user.ID))
	return err
}

// DeleteProfile delete profile
func (f *Facade) DeleteProfile(ctx context.Context, req *proto.UpdateProfileRequest, rsp *proto.UpdateProfileResponse) error {
	pid, email := req.GetPid(), req.GetEmail()
	user, err := f.parseUser(pid, email)
	if err != nil {
		return err
	}

	profilecontroller := controllers.NewProfileController()
	err = profilecontroller.DeleteProfile(&models.Profile{
		Year:      req.GetYear(),
		Class:     req.GetClass(),
		Name:      req.GetName(),
		Date:      req.GetDate(),
		CreatedBy: user.Email,
	})

	f.syslog(notification.GrowthProfileUpdated(user.ID))
	return err
}
