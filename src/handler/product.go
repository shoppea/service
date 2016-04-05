package handler

import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
	"github.com/Sirupsen/logrus"
)

func FindProduct(c* gin.Context)  {
	//var products []product.Product
	//name := c.Param("name")
	//common.GetAllRows(products,name)
}

func CreateProduct(c *gin.Context) {
	var product product.Product
	common.BindResponse(c,&product)
	logrus.Info("%+v",product)
	product.Add(c)
	//err := product.Add()
	//if err != nil {
	//	throw.ErrorDB(c,err)
	//}
}