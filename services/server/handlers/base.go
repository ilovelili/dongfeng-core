// Package handlers define the core behaviors of each API
package handlers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	notification "github.com/ilovelili/dongfeng-notification"
	sharedlib "github.com/ilovelili/dongfeng-shared-lib"
)

// Facade api facade
type Facade struct {
	AuthClient *sharedlib.Client
}

// NewFacade constructor
func NewFacade() *Facade {
	config := utils.GetConfig()
	return &Facade{
		AuthClient: sharedlib.NewAuthClient(config.Auth.ClientID, config.Auth.ClientSecret),
	}
}

// syslog save notification
func (f *Facade) syslog(notification *notification.Notification) {
	go func() {
		notificationcontroller := controllers.NewNotificationController()
		notificationcontroller.Save([]*models.Notification{
			&models.Notification{
				UserID:     notification.UserID,
				CustomCode: notification.CustomCode,
				Details:    notification.Details,
				Link:       notification.Link,
				CategoryID: notification.CategoryID,
				Read:       0,
				Time:       notification.Time,
			},
		})
	}()
}
