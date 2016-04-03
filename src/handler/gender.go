package handler

import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
)

func CreateGender(c *gin.Context) {
	var g product.Gender
	common.BindResponse(c,&g)
	common.InsertDB(g)
}
