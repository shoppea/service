package handler

import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
	"db"
	"throw"
)

func CreateCategory(c *gin.Context)  {
	var category product.Category
	common.BindResponse(c,&category)
	common.InsertDB(category)
}

func CreateSubCategory(c *gin.Context) {
	var subCategory product.SubCategory
	common.BindResponse(c,&subCategory)
	common.InsertDB(subCategory)
}

func GetAllCategories(c *gin.Context) {
	var categories []product.Category
	dbPool := db.SC()
	err := dbPool.Find(&categories).Error
	if err != nil {
		throw.ErrorDB(c,err)
	}else {
		throw.SuccessOK(c,categories)
	}
}