package common

import "github.com/gin-gonic/gin"

func GetTestModeGinRouter() *gin.Engine {
	gin.SetMode(gin.TestMode);
	return gin.Default();
}