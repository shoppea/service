package coming_soon

import (
	"db"
)

// User struct which will get
// retrieved by user from
// coming_soon page
type User struct {
	Email string
	//Date time.UnixDate
}

// Register this user to
// database
func (u *User ) RegisterForNotify() error {
	dbPool := db.SharedConnection()
	return dbPool.Create(u).Error
}
