package throw

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorBadRequest(c *gin.Context,err error) {
	c.JSON(http.StatusBadRequest,gin.H{
		"message": err.Error(),
	})
	c.Abort()
}

func ErrorDB(c *gin.Context,err error) {
	c.JSON(http.StatusBadRequest,gin.H{
		"error": err.Error(),
	})
	c.Abort()
}
