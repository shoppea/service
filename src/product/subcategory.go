package product

type SubCategory struct {
	Id int
	Name string
	Category Category
	CategoryID int      `json:"category_id"`
	Gender Gender
	GenderID int        `json:"gender_id"`
}