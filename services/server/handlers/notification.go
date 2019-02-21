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

// SaveNotification save notification
func (f *Facade) SaveNotification(ctx context.Context, req *proto.SaveNotificationRequest, rsp *proto.SaveNotificationResponse) error {
	notification := req.GetNotification()
	notificationcontroller := controllers.NewNotificationController()

	return notificationcontroller.Save([]*models.Notification{
		&models.Notification{
			UserID:     notification.GetUserId(),
			CustomCode: notification.GetCustomCode(),
			Details:    notification.GetDetails(),
			Link:       notification.GetLink(),
			CategoryID: notification.GetCategoryId(),
			Read:       notification.GetRead(),
		},
	})
}

// UpdateNotification update notification
func (f *Facade) UpdateNotification(ctx context.Context, req *proto.UpdateNotificationsRequest, rsp *proto.UpdateNotificationsResponse) error {
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
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	ids := req.GetNotifications()
	notificationcontroller := controllers.NewNotificationController()
	notifications := make([]*models.Notification, 0)
	for _, id := range ids {
		notifications = append(notifications, &models.Notification{
			ID: id,
		})
	}

	return notificationcontroller.Save(notifications)
}
