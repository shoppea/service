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
	Email string
	CreatedAt time.Time
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
