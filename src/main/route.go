package main

import (
	"db"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"handler"
	"net/http" 
)

func init() {
	dbPool := db.SharedConnection()
	if dbPool.Error != nil {
		logrus.Error(dbPool.Error)
		panic("failed to initialse database")
	}
}

func main() {
	r := gin.Default()

	r.StaticFile("/swagger_api","../src/main/swagger.json")

	r.StaticFS("/docs",http.Dir("../dist"))

	r.POST("/category", handler.CreateCategory)
	r.GET("/categories", handler.GetAllCategories)

	r.POST("/gender", handler.CreateGender)
	r.GET("/genders", handler.GetAllGenders)

	r.POST("/sub_category", handler.CreateSubCategory)

	r.POST("/product", handler.CreateProduct)
	r.GET("/products", handler.GetAllProducts)

	r.POST("/user", handler.CreateUser)

	r.GET("/help", func(c *gin.Context) {
		c.JSON(http.StatusOK, r.Routes())
	})

	r.GET("/search", handler.FindProduct)

	r.Run(":8081")
}