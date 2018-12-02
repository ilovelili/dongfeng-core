// Package handlers define the core behaviors of each API
package handlers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	notification "github.com/ilovelili/dongfeng-notification"
)

// Facade api facade
type Facade struct{}

// NewFacade constructor
func NewFacade() *Facade {
	return new(Facade)
}

// syslog save notification
func (f *Facade) syslog(notification *notification.Notification) {
	go func() {
		notificationcontroller := controllers.NewNotificationController()
		notificationcontroller.Save(&models.Notification{
			UserID:     notification.UserID,
			CustomCode: notification.CustomCode,
			Details:    notification.Details,
			Link:       notification.Link,
			CategoryID: notification.CategoryID,
			Time:       notification.Time,
		})
	}()
}
