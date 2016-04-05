package main

import (
	"db"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"handler"
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

	//r.GET("/product/:name",handler.FindProduct)
	r.POST("/category",handler.CreateCategory)

	r.POST("/gender",handler.CreateGender)

	r.POST("/sub_category",handler.CreateSubCategory)

	r.POST("/product",handler.CreateProduct)

	r.Run(":8080")
}