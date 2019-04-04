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

// GetClasses get classes
func (f *Facade) GetClasses(ctx context.Context, req *proto.GetClassRequest, rsp *proto.GetClassResponse) error {
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
