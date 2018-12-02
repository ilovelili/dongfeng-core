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

// UpdateUser update user info
func (f *Facade) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest, rsp *proto.UpdateUserResponse) error {
	if req.GetName() == "" && req.GetAvatar() == "" {
		return utils.NewError(errorcode.CoreInvalidUpdateUserRequest)
	}

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

	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	if req.GetName() != "" {
		user.Name = req.GetName()
	}

	if req.GetAvatar() != "" {
		user.Avatar = req.GetAvatar()
	}

	err = usercontroller.Save(user)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToSaveUser)
	}

	f.syslog(notification.ProfileUpdated(user.ID))
	return err
}
