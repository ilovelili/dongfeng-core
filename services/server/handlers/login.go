// Package handlers define the core behaviors of each API
package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-errorcode"
	proto "github.com/ilovelili/dongfeng-protobuf"
	"github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

const (
	// he is agent smith
	agentsmith = "AgentSmith"
)

// Login return 200
func (f *Facade) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return utils.NewError(errorcode.GenericInvalidMetaData)
	}

	idtoken := md[sharedlib.MetaDataToken]
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
	var newUser bool
	usercontroller := controllers.NewUserController()
	exsitinguser, err := usercontroller.GetUserByEmail(user.Email)
	newUser = err != nil

	if newUser {
		if user.Role == "" {
			user.Role = "user"
		}

		// new user, save to database
		if err = usercontroller.Save(user); err != nil {
			return utils.NewError(errorcode.CoreFailedToSaveUser)
		}
	} else {
		user = exsitinguser
	}

	// set version
	rsp.Version = utils.Version()

	// user profile
	rsp.User = &proto.User{
		Newuser:  newUser,
		Role:     user.Role,
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Settings: convertSettings(user.Settings),
	}

	// fetch operation logs
	operationcontroller := controllers.NewOperationController()

	// 1. system broadcasting
	uid := agentsmith
	operations, err := operationcontroller.GetOperations(uid, true)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetOperation)
	}
	rsp.SystemBroadcasting = convertOperations(operations)

	// 2. user updates
	operations, err = operationcontroller.GetOperations(user.ID, user.Role == "admin")
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetOperation)
	}
	rsp.UserUpdates = convertOperations(operations)

	// 3. friends updates
	friendcontroller := controllers.NewFriendController()
	friends, err := friendcontroller.GetFriends(user.ID)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetFriends)
	}

	operations = make([]*models.Operation, 0)
	for _, friend := range friends {
		ops, err := operationcontroller.GetOperations(friend.ID, friend.Role == "admin")
		if err != nil {
			return utils.NewError(errorcode.CoreFailedToGetOperation)
		}

		operations = append(operations, ops...)
	}
	rsp.FriendUpdates = convertOperations(operations)

	return err
}

func convertSettings(settings []*models.Settings) []*proto.Setting {
	result := make([]*proto.Setting, 0)
	for _, setting := range settings {
		result = append(result, &proto.Setting{
			Id:      setting.ID,
			Name:    setting.Name,
			Enabled: setting.Enabled,
		})
	}

	return result
}

func convertOperations(operations []*models.Operation) []*proto.Operation {
	result := make([]*proto.Operation, 0)
	for _, operation := range operations {
		var uid string
		if operation.UserID != agentsmith {
			uid = operation.UserID
		}

		result = append(result, &proto.Operation{
			UserId:    uid,
			Time:      operation.Time,
			Operation: operation.Operation,
			Category:  operation.Category,
		})
	}

	return result
}
