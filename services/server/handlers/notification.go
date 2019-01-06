package handlers

import (
	"context"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	proto "github.com/ilovelili/dongfeng-protobuf"
)

// SaveNotification save notification
func (f *Facade) SaveNotification(ctx context.Context, req *proto.SaveNotificationRequest, rsp *proto.SaveNotificationResponse) error {
	notification := req.GetNotification()
	notificationcontroller := controllers.NewNotificationController()

	return notificationcontroller.Save(&models.Notification{
		UserID:     notification.GetUserId(),
		CustomCode: notification.GetCustomCode(),
		Details:    notification.GetDetails(),
		Link:       notification.GetLink(),
		CategoryID: notification.GetCategoryId(),
	})
}
