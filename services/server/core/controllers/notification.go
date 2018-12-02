package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// NotificationController user profile controller
type NotificationController struct {
	repository *repositories.NotificationRepository
}

// NewNotificationController new controller
func NewNotificationController() *NotificationController {
	return &NotificationController{
		repository: repositories.NewNotificationRepository(),
	}
}

// GetNotifications get Notifications
func (c *NotificationController) GetNotifications(uid string, adminonly bool) ([]*models.Notification, error) {
	return c.repository.SelectNotifications(uid, adminonly)
}

// Save save Notifications
func (c *NotificationController) Save(notification *models.Notification) error {
	return c.repository.Insert(notification)
}
