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

// GetPhysiques get pupils
func (f *Facade) GetPhysiques(ctx context.Context, req *proto.GetPhysiqueRequest, rsp *proto.GetPhysiqueResponse) error {
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

	return nil
}

// UpdatePhysique update physique
func (f *Facade) UpdatePhysique(ctx context.Context, req *proto.UpdatePhysiqueRequest, rsp *proto.UpdatePhysiqueResponse) error {
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

	physiques := req.GetPhysiques()
	if len(physiques) != 1 {
		return utils.NewError(errorcode.CoreFailedToUpdatePhysiques)
	}

	physique := physiques[0]
	physique.CreatedBy = exsitinguser.Email

	// physiquescontroller := controllers.NewPhysiqueController()
	// err = pupilcontroller.UpdatePhysique(&models.Physique{
	// 	ID:    pupil.GetId(),
	// 	Name:  pupil.GetName(),
	// 	Class: pupil.GetClass(),
	// })

	f.syslog(notification.PhysiqueUpdated(exsitinguser.ID))
	return err
}

// UpdatePhysiques update pupils
func (f *Facade) UpdatePhysiques(ctx context.Context, req *proto.UpdatePhysiqueRequest, rsp *proto.UpdatePhysiqueResponse) error {
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

	physiques := req.GetPhysiques()
	for _, physique := range physiques {
		physique.CreatedBy = exsitinguser.Email
	}

	// pupilcontroller := controllers.NewPhysiqueController()
	// err = pupilcontroller.UpdatePhysiques(pupils)
	// if err != nil {
	// 	return utils.NewError(errorcode.CoreFailedToUpdatePhysiques)
	// }

	f.syslog(notification.NamelistUpdated(exsitinguser.ID))
	return nil
}
