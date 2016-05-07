package handler

import (
	"github.com/gin-gonic/gin"
	"user"
	"common"
	"throw"
	"db"
	"strconv"
)

func AddToCart(c *gin.Context)  {
	var cart user.Cart
	var err error
	err = common.BindResponse(c,&cart)
	if err != nil {
		throw.ErrorBadRequest(c,err)
		return
	}
	dbPool := db.SC()
	err = dbPool.Create(&cart).Error
	if err != nil {
		throw.ErrorDB(c,err)
	}else{
		throw.SuccessCreated(c,cart)
		return
	}
}

func GetUserCart(c *gin.Context) {
	var cart []user.Cart
	var err error
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		throw.ErrorBadRequest(c,err)
		return
	}
	user := user.User{Id:id}
	dbPool := db.SC()
	err = dbPool.Model(&user).Related(&cart).Error
	if err != nil {
		throw.ErrorDB(c,err)
		return
	}
	common.AfterService(c,err,cart)
}