package common

import (
	"github.com/gin-gonic/gin"
	"throw"
)

func AfterService(c *gin.Context, err error, obj interface{}) {
	if err != nil {
		throw.ErrorBadRequest(c,err)
		return
	}else{
		throw.SuccessOK(c,obj)
		return
	}
}