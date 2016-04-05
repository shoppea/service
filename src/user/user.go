package user

import "product"

type IUser interface {
	Add()
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
	IUser
	UFirstName string
	ULastName string
	UAddress Address
	UContact int
}

func (u *User ) Add() {
	// Adding user to db
}