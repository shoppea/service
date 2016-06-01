package coming_soon

import (
	"db"
	"time"
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
	return dbPool.Create(u).Error
}
