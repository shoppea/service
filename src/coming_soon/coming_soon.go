package coming_soon

import (
	"db"
	"time"
	"errors"
)

// User struct which will get
// retrieved by user from
// coming_soon page
type NotifyUser struct {
	Email string        `json:"email"`
	CreatedAt time.Time        `json:"created_at"`
}

// Register this user to
// database
func (u *NotifyUser) RegisterForNotify() error {
	dbPool := db.SharedConnection()
	u.CreatedAt = time.Now()
	if u.Email == "" {
		return errors.New("Bad gateway")
	}
	return dbPool.Create(u).Error
}

// This API Call returns all registered
// users
func GetAllRegisteredUsers() (users []NotifyUser,err error) {
	dbPool := db.SharedConnection()
	err = dbPool.Find(&users).Error
	return users,err
}