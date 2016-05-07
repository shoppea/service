package user

import (
	"product"
	"db"
)

type IUser interface {
	Add() (err error)
	AddToCart(product *product.Product)
}

type Address struct {
	HouseName string
	StreetNo string
	AreaName string
	City string
	PinCode int
}

type User struct {
	IUser			`gorm:"-"`
	Id int                  `json:"id"`
	FirstName string        `json:"first_name"`
	LastName string 	`json:"last_name"`
	ContactNo string        `json:"contact_no"`
	Email string        	`json:"email"`
	Password string         `json:"password"`
	Carts []Cart        	`json:"carts"`
}

func (u *User ) Add() (err error) {
	dbPool := db.GetConnection()
	return dbPool.Create(u).Error
}