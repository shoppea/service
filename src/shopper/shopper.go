package shopper

import (
	"common"
	"github.com/gin-gonic/gin"
)

type IShopper interface{
	Add()
}

// swagger:response ShopperCreate
type Shopper struct {
	common.BaseService	`gorm:"-"`
	Id int                `json:"id"`
	Name string        `json:"name"`
	ShopOwner string `json:"owner_name"`
	Description string        `json:"description"`
	Address string        `json:"address"`
	VatNo string        `json:"vat_no"`
	ContactNo int        `json:"contact_no"`
	Latitude string        `json:"latitude"`
	Longitude string        `json:"longitude"`
	AreaName string        `json:"area_name"`
	//Products []product.Product        `gorm:"many2many:shopper_stock"`
}

func New() *Shopper {
	return &Shopper{}
}

func (s *Shopper ) AddShopper(c *gin.Context) error {
	 return common.InsertDBWithContext(c,s)
}

type ShopperStock struct {
	Id int
	ProductID int        `json:"product_id"`
	ShopperID int        `json:"shopper_id"`
	Quantity int         `json:"quantity"`
	Price float64        `json:"price"`
}
