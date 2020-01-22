package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
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
	pid, email := req.GetPid(), req.GetEmail()
	_, err := f.parseUser(pid, email)
	if err != nil {
		return err
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
