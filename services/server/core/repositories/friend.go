package repositories

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// FriendRepository friends repository
type FriendRepository struct{}

// NewFriendRepository init UserProfile repository
func NewFriendRepository() *FriendRepository {
	return &FriendRepository{}
}

// Select select friends by user id
func (r *FriendRepository) Select(uid string) (friends []*models.User, err error) {
	query := Table("user_friends").Alias("f").
		Join("users").Alias("u").On("f.friend_id", "u.id").
		Project("u.*").
		Where().Eq("f.user_id", uid).Sql()

	// no rows is actually not an error
	if err = session().Find(query, nil).All(&friends); err != nil && norows(err) {
		err = nil
	}

	return
}

// Insert insert user and her friend
func (r *FriendRepository) Insert(userfriend *models.UserFriend) error {
	query := Table("user_friends").Alias("f").
		Project(`count("f.*")`).
		Where().Eq("f.user_id", userfriend.UserID).Eq("f.friend_id", userfriend.FriendID).Sql()

	count, err := session().Count(query, nil)
	if err != nil || 0 == count {
		return insertTx(userfriend)
	}

	return nil
}
