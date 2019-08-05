package handlers

import (
	"context"
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
	proto "github.com/ilovelili/dongfeng-protobuf"	
)

// UpdateNotification update notification
func (f *Facade) UpdateNotification(ctx context.Context, req *proto.UpdateNotificationRequest, rsp *proto.UpdateNotificationResponse) error {
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

// UpdateNotifications update notification
func (f *Facade) UpdateNotifications(ctx context.Context, req *proto.UpdateNotificationsRequest, rsp *proto.UpdateNotificationsResponse) error {
	pid := req.GetPid()
	userinfo, err := f.AuthClient.ParseUserInfo(pid)	
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	var user *models.User
	err = json.Unmarshal(userinfo, &user)
	if err != nil {
		return utils.NewError(errorcode.GenericInvalidToken)
	}

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		return utils.NewError(errorcode.CoreNoUser)
	}

	ids := req.GetNotifications()
	notificationcontroller := controllers.NewNotificationController()
	notifications := []*models.Notification{}
	for _, id := range ids {
		notifications = append(notifications, &models.Notification{
			ID: id,
		})
	}

	return notificationcontroller.Save(notifications)
}
