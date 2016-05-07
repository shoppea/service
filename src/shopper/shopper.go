package shopper

import (
	"product"
)

type IShopper interface{
	Add()
}

type Shopper struct {
	//gorm.Model
	Id int
	Name string
	Description string
	//SAddress user.Address
	//VatNo string
	//SContactNo int
	Products []product.Product        `gorm:"many2many:shopper_stock"`
}

func New() *Shopper {
	return &Shopper{}
}

type ShopperStock struct {
	Id int
	ProductID int        `json:"product_id"`
	ShopperID int        `json:"shopper_id"`
	Quantity int         `json:"quantity"`
	Price float64        `json:"price"`
}

