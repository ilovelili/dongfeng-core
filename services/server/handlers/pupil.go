package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-error-code"
	notification "github.com/ilovelili/dongfeng-notification"
	proto "github.com/ilovelili/dongfeng-protobuf"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// GetPupils get pupils
func (f *Facade) GetPupils(ctx context.Context, req *proto.GetPupilRequest, rsp *proto.GetPupilResponse) error {
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

	pupils := req.GetPupils()
	if len(pupils) != 1 {
		return utils.NewError(errorcode.CoreFailedToUpdatePupils)
	}

	pupil := pupils[0]
	pupil.CreatedBy = exsitinguser.Email

	pupilcontroller := controllers.NewPupilController()
	err = pupilcontroller.UpdatePupil(&models.Pupil{
		ID:    pupil.GetId(),
		Name:  pupil.GetName(),
		Class: pupil.GetClass(),
	})

	f.syslog(notification.NamelistUpdated(exsitinguser.ID))
	return err
}

// UpdatePupils update pupils
func (f *Facade) UpdatePupils(ctx context.Context, req *proto.UpdatePupilRequest, rsp *proto.UpdatePupilResponse) error {
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
