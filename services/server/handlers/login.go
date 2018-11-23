// Package handlers define the core behaviors of each API
package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	"github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"
	"github.com/ilovelili/dongfeng-shared-lib"
	"github.com/micro/go-micro/metadata"
)

// Login handler returns all data needed by front end
func (f *Facade) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
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
		Settings: resolveSettings(user.Settings),
	}

	// fetch operation logs
	notificationcontroller := controllers.NewNotificationController()
	notifications, err := notificationcontroller.GetNotifications(user.ID, true)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetNotification)
	}
	rsp.Notifications = resolvecNotifications(notifications)
	return err
}

func resolveSettings(settings []*models.Settings) []*proto.Setting {
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

func resolvecNotifications(notifications []*models.Notification) []*proto.Notification {
	result := make([]*proto.Notification, 0)
	for _, notification := range notifications {
		result = append(result, &proto.Notification{
			UserId:     notification.UserID,
			CustomCode: notification.CustomCode,
			Category:   notification.Category,
			Details:    notification.Details,
			Link:       notification.Link,
			Time:       notification.Time,
		})
	}

	return result
}
