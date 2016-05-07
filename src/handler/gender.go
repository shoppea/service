package handler

import (
	"github.com/gin-gonic/gin"
	"product"
	"common"
	"db"
	"throw"
)

func CreateGender(c *gin.Context) {
	var g product.Gender
	common.BindResponse(c,&g)
	common.InsertDB(g)
}

func GetAllGenders(c *gin.Context) {
	var genders []product.Gender
	dbPool := db.SC()
	err := dbPool.Find(&genders).Error
	if err != nil {
		throw.ErrorDB(c,err)
	}else {
		throw.SuccessOK(c,genders)
	}
}
