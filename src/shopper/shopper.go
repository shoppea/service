package shopper

import (
	"user"
	"github.com/jinzhu/gorm"
)

type IShopper interface{
	Add()
}

type Shopper struct {
	gorm.Model
	SName string
	SAddress user.Address
	VatNo string
	SContactNo int
}