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

// GetClasslist get class list master data
func (f *Facade) GetClasslist(ctx context.Context, req *proto.GetClasslistRequest, rsp *proto.GetClasslistResponse) error {
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

	items := make([]*proto.ClassItem, 0)
	for _, class := range classes {
		items = append(items, &proto.ClassItem{
			Name:      class.Name,
			Id:        class.ID,
			CreatedBy: class.CreatedBy,
		})
	}

	rsp.Items = items
	return nil
}

// UpdateClasslist update class list master data
func (f *Facade) UpdateClasslist(ctx context.Context, req *proto.UpdateClasslistRequest, rsp *proto.UpdateClasslistResponse) error {
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

	classes := req.GetItems()
	for _, class := range classes {
		class.CreatedBy = exsitinguser.Email
	}

	classcontroller := controllers.NewClassController()
	err = classcontroller.UpdateClasses(classes)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToUpdateClasses)
	}

	f.syslog(notification.ClasslistUpdated(user.ID))
	return nil
}
