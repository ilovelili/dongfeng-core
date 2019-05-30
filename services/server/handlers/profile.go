package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetProfile get profile
func (f *Facade) GetProfile(ctx context.Context, req *proto.GetProfileRequest, rsp *proto.GetProfileResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	_, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
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

// GetProfiles get profile
func (f *Facade) GetProfiles(ctx context.Context, req *proto.GetProfilesRequest, rsp *proto.GetProfilesResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	_, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
		return utils.NewError(errorcode.GenericInvalidToken)
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
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := req.GetToken()
	jwks := md[sharedlib.MetaDataJwks]
	claims, token, err := sharedlib.ParseJWT(idtoken, jwks)

	// vaidate the token
	if err != nil || !token.Valid {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// Unmarshal user info
	userinfo, _ := json.Marshal(claims)
	var user *models.User
	err = json.Unmarshal(userinfo, &user)

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	profilecontroller := controllers.NewProfileController()
	err = profilecontroller.SaveProfile(&models.Profile{
		Year:      req.GetYear(),
		Class:     req.GetClass(),
		Name:      req.GetName(),
		Date:      req.GetDate(),
		Profile:   req.GetProfile(),
		Enabled:   req.GetEnabled(),
		CreatedBy: exsitinguser.Email,
	})

	return err
}
