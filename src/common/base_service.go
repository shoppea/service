package common

import (
	"github.com/gin-gonic/gin"
)

type BaseService struct {

}

func (b *BaseService ) Add(c *gin.Context) error {
	return InsertDBWithContext(c,b)
}