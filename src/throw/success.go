package throw

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessCreated(c *gin.Context,obj interface{})  {
	c.JSON(http.StatusCreated,obj)
}

func SuccessOK(c *gin.Context,obj interface{})  {
	c.JSON(http.StatusOK,obj)
}