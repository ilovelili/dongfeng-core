package repositories

import (
	"fmt"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// NotificationRepository friends repository
type NotificationRepository struct{}

// NewNotificationRepository init UserProfile repository
func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{}
}

// Select select notification logs
func (r *NotificationRepository) Select(uid string, adminonly bool) (notifications []*models.Notification, err error) {
	query := fmt.Sprintf("CALL spSelectNotifications('%s', %d)", uid, resolveAdminOnly(adminonly))
	err = session().Find(query, nil).All(&notifications)
	if norows(err) {
		err = nil
	}

	return
}

// Insert insert Notification
func (r *NotificationRepository) Insert(Notification *models.Notification) error {
	return insertTx(Notification)
}

func resolveAdminOnly(adminonly bool) int {
	if adminonly {
		return 1
	}
	return 0
}
