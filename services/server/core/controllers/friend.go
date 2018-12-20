package controllers

import (
	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
)

// FriendController user profile controller
type FriendController struct {
	repository *repositories.FriendRepository
}

// NewFriendController new controller
func NewFriendController() *FriendController {
	return &FriendController{
		repository: repositories.NewFriendRepository(),
	}
}

// GetFriends get user friends
func (c *FriendController) GetFriends(uid string) ([]*models.User, error) {
	return c.repository.Select(uid)
}

// Save save friends
func (c *FriendController) Save(userfriend *models.UserFriend) error {
	return c.repository.Insert(userfriend)
}
