package handler

import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
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