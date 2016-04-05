package handler

import (
	"github.com/gin-gonic/gin"
	"common"
	"user"
	"throw"
)

func CreateUser(c *gin.Context) {
	var user user.User
	err := common.BindResponse(c,&user)
	if err != nil {
		throw.ErrorDB(c,err)
		c.Abort()
	}
	err = user.Add()
	common.AfterService(c,err,user)
}