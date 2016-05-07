package db


import (
	_ "github.com/go-sql-driver/mysql"
	_ "sync"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/Sirupsen/logrus"
	"sync"
)

var DBConnection *gorm.DB
var once sync.Once

// Abbr. Shared Connection cool to connect with mysql database
// GoLang DB instance automatically handles database pooling

func SharedConnection() ( *gorm.DB ){
	once.Do(func() {
		var err error
		DBConnection,err = gorm.Open("mysql", "root:root@tcp(ec2-54-164-31-179.compute-1.amazonaws.com:3306)/snabar_staging")
		DBConnection.LogMode(true)
		if err != nil {
			logrus.Info("Error encoutered while getting connection from mysql")
			fmt.Println(err)
		}
		logrus.Info("Getting connection from MYSQL")
		return
	})
	return DBConnection
}

func GetConnection() *gorm.DB {
	pool := SharedConnection()
	return pool
}