package handler

import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
)

func FindProduct(c* gin.Context)  {
	//var products []product.Product
	//name := c.Param("name")
	//common.GetAllRows(products,name)
}

func CreateProduct(c *gin.Context) {
	var product product.Product
	common.BindResponse(c,product)
}