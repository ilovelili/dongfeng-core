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
)

// Dashboard handler returns data needed by dashboard
func (f *Facade) Dashboard(ctx context.Context, req *proto.DashboardRequest, rsp *proto.DashboardResponse) error {
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

	rsp.UserId = user.ID
	// fetch operation logs
	notificationcontroller := controllers.NewNotificationController()
	notifications, err := notificationcontroller.GetNotifications(user.ID, true, true)
	if err != nil {
		return utils.NewError(errorcode.CoreFailedToGetNotification)
	}
	rsp.Notifications = resolvecNotifications(notifications)
	return err
}

func resolvecNotifications(notifications []*models.Notification) []*proto.Notification {
	result := []*proto.Notification{}
	for _, notification := range notifications {
		result = append(result, &proto.Notification{
			Id:         int64(notification.ID),
			UserId:     notification.UserID,
			CustomCode: notification.CustomCode,
			Category:   notification.Category(),
			CategoryId: notification.CategoryID,
			Details:    notification.Details,
			Link:       notification.Link,
			Time:       sharedlib.NewTime(notification.Time).FormatTime(),
		})
	}

	return result
}
