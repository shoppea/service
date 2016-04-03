package common

import (
	"github.com/gin-gonic/gin"
	"db"
)

func BindResponse(c *gin.Context,obj interface{}) (err error) {
	err = c.BindJSON(&obj)
	return err
}

func InsertDB(obj interface{}) (err error) {
	dbPool := db.SharedConnection()

	err = dbPool.Error
	if err != nil {
		return
	}

	err = dbPool.Create(obj).Error
	return
}

//func GetAllRows(obj []interface{},name string) {
//	dbPool := db.SharedConnection()
//	dbPool.Where("name=?",name).Find(&obj)
//}
