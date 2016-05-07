package common

import (
	"github.com/gin-gonic/gin"
	"db"
	"throw"
	"github.com/Sirupsen/logrus"
)

func BindResponse(c *gin.Context,obj interface{}) (err error) {
	err = c.BindJSON(obj)
	logrus.Info("%#v",obj)
	if err != nil {
		throw.ErrorDB(c,err)
		c.Abort()
	}
	return err
}

func InsertDataBase(c *gin.Context,obj interface{}) (err error) {
	dbPool := db.SharedConnection()
	err = dbPool.Error
	if err != nil {
		throw.ErrorDB(c,err)
		return
	}
	err = dbPool.Create(obj).Error
	return
}


func InsertDBWithContext(c *gin.Context,obj interface{}) (err error) {
	dbPool := db.SharedConnection()
	err = dbPool.Error
	if err != nil {
		throw.ErrorDB(c,err)
	}
	logrus.Info("%+v",obj)
	err = dbPool.Create(obj).Error
	if err != nil {
		throw.ErrorDB(c,err)
	}
	return
}

func InsertDB(obj interface{}) (err error) {
	dbPool := db.SharedConnection()
	err = dbPool.Error
	if err != nil {
		return
	}
	logrus.Info("%+v",obj)
	err = dbPool.Create(obj).Error
	return
}

//func GetAllRows(obj []interface{},name string) {
//	dbPool := db.SharedConnection()
//	dbPool.Where("name=?",name).Find(&obj)
//}
