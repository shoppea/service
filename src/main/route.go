package main

import (
	"db"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"handler"
	"net/http"
)

//go:generate swagger generate spec

func init() {
	dbPool := db.SC()
	if dbPool.Error != nil {
		logrus.Error(dbPool.Error)
		panic("failed to initialse database")
	}
}

func main() {

	r := gin.Default()

	// swagger:route GET /pets listPets pets users
	//
	// Lists pets filtered by some parameters.
	//
	// This will show all available pets by default.
	// You can get the pets that are out of stock
	//
	//     Consumes:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Produces:
	//     - application/json
	//     - application/x-protobuf
	//
	//     Schemes: http, https, ws, wss
	//
	//     Security:
	//       api_key:
	//       oauth: read, write
	//
	//     Responses:

	// Auth accepts two params
	r.POST("/auth",handler.Authenticate)

	r.POST("/category",handler.CreateCategory)
	r.GET("/categories",handler.GetAllCategories)

	r.POST("/gender",handler.CreateGender)
	r.GET("/genders",handler.GetAllGenders)

	r.POST("/sub_category",handler.CreateSubCategory)

	r.POST("/product",handler.CreateProduct)
	r.GET("/products",handler.GetAllProducts)

	r.POST("/shop", handler.AddShopper)
	r.POST("/stock/add", handler.AddStockToShopper)

	// User related handlers
	r.POST("/user",handler.CreateUser)
	r.POST("/cart",handler.AddToCart)
	r.GET("/cart/:id",handler.GetUserCart)

	r.GET("/help", func(c *gin.Context) {
		c.JSON(http.StatusOK,r.Routes())
	})

	r.Run(":8080")
}