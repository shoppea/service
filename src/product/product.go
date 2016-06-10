package product

import (
	"common"
	"github.com/gin-gonic/gin"
)

type IProduct interface {
	Add(c *gin.Context) (err error)
	AddStock(quantity int)
}

type Product struct {
	IProduct
	Id int        `json:"id"`
	Name string     `json:"product_name"`
	Price int        `json:"price"`
	Weight float64  `json:"product_weight"`
	ShortDesc string `json:"product_short_description"`
	LongDesc string        `json:"product_long_description"`
	CategoryID int                `json:"category_id"`
	SubCategory SubCategory
	Image string                `json:"image"`
}

func (p *Product) Add(c *gin.Context) (err error){
	return common.InsertDBWithContext(c,p)
}
