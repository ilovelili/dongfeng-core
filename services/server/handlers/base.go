// Package handlers define the core behaviors of each API
package handlers

import (
	"encoding/json"

	"github.com/ilovelili/dongfeng-core/services/server/core/controllers"
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/utils"
	errorcode "github.com/ilovelili/dongfeng-error-code"
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

// parseUser parse user
func (f *Facade) parseUser(pid, email string) (user *models.User, err error) {
	userinfo, err := f.AuthClient.ParseUserInfo(pid)
	if err != nil {
		err = utils.NewError(errorcode.GenericInvalidToken)
		return
	}

	err = json.Unmarshal(userinfo, &user)
	if err != nil {
		err = utils.NewError(errorcode.GenericInvalidToken)
		return
	}

	// if user email is empty (for example, open id login without email bound), set it
	if user.Email == "" {
		user.Email = email
	}

	// check if user exists or not
	usercontroller := controllers.NewUserController()
	user, err = usercontroller.GetUserByEmail(user.Email)
	if err != nil {
		err = utils.NewError(errorcode.CoreNoUser)
	}

	return
}
