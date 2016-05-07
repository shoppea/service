package handler

import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
	"github.com/Sirupsen/logrus"
	"throw"
	"db"
)

func FindProduct(c*gin.Context) {
	var products []product.Product
	searchStr := c.Query("query")
	if len(searchStr) != 0 {
		dbPool := db.SharedConnection()
		err := dbPool.Where("name LIKE ?", "%" + c.Query("query") + "%").Find(&products).Error
		if err != nil {
			throw.ErrorDB(c, err)
		} else {
			throw.SuccessOK(c, products);
		}
	} else {
		throw.SuccessOK(c, products);
	}
}

func CreateProduct(c *gin.Context) {
	var product product.Product
	err := common.BindResponse(c, &product)
	if err != nil {
		throw.ErrorBadRequest(c, err)
		return
	}
	logrus.Info("%+v", product)
	err = product.Add(c)
	if err == nil {
		throw.SuccessCreated(c, product)
		return
	} else {
		throw.ErrorDB(c, err)
		return
	}
}

func GetAllProducts(c *gin.Context) {
	var products []product.Product
	dbPool := db.SharedConnection()
	err := dbPool.Find(&products).Error
	if err != nil {
		throw.ErrorDB(c, err)
	} else {
		throw.SuccessOK(c, products)
	}
}