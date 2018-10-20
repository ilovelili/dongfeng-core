package controllers

import (
	"math/rand"
	"time"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
	"github.com/ilovelili/dongfeng-core/services/server/core/repositories"
	"github.com/ilovelili/dongfeng-shared-lib"
)

// UserController user profile controller
type UserController struct {
	repository *repositories.UserRepository
}

// NewUserController new controller
func NewUserController() *UserController {
	return &UserController{
		repository: repositories.NewUserRepository(),
	}
}

// GetUserByEmail get user by mail
func (c *UserController) GetUserByEmail(email string) (*models.User, error) {
	return c.repository.SelectByMail(email)
}

// Save save user
func (c *UserController) Save(user *models.User) error {
	if user.ID == "" {
		for uid, exist := c.generateUID(); ; {
			// conflict, regenerate
			if exist {
				// reset random seed
				time.Sleep(time.Duration(rand.Intn(1e2)) * time.Millisecond)
			} else {
				user.ID = uid
				break
			}
		}
	}

	return c.repository.Upsert(user)
}

func (c *UserController) generateUID() (string, bool) {
	uid := sharedlib.RandStringRunes(12)
	count, _ := c.repository.CountByID(uid)
	return uid, count > 0
}
