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

	r.Run(":8080")
}