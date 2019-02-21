package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// NotificationController notification controller
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
func (c *NotificationController) GetNotifications(uid string, adminonly bool, excluderead bool) ([]*models.Notification, error) {
	return c.repository.Select(uid, adminonly, excluderead)
}

// Save save Notifications
func (c *NotificationController) Save(notifications []*models.Notification) error {
	return c.repository.Upsert(notifications)
}
