package electronics

import "product"

type ElectronicProduct struct {
	Id int
	Name string
	Status string
	Price float64
	Description string
	SubcategoryID int			`json:"subcategory_id"`
	SubCategory product.SubCategory
}
