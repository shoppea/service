package handler

import (
	"github.com/gin-gonic/gin"
	"common"
	"user"
	"throw"
	"db"
	"github.com/Sirupsen/logrus"
)

type UserAuth struct{
	Authenticated bool
}

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

func Authenticate(c *gin.Context)  {
	var user user.User
	var auth UserAuth
	var err error
	var creds struct{
		UserName string        `json:"user_name"`
		Password string
	}
	err = common.BindResponse(c,&creds)
	if err != nil {
		throw.ErrorBadRequest(c,err)
		return
	}
	dbPool := db.SharedConnection()
	err = dbPool.Where("email = ? AND password = ?", creds.UserName,creds.Password).Find(&user).Error
	logrus.Info("User is : %+v", user.Id)
	if user.Id > 0 {
		auth.Authenticated = true
	}else {
		c.Status(404)
		auth.Authenticated = false
	}
	throw.SuccessOK(c,auth)
}