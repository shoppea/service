package product

type IGender interface {

}

type Gender struct {
	IGender
	Id int
	Name string
}

