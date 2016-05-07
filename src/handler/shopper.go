package handler

import (
	"github.com/gin-gonic/gin"
	"shopper"
	"common"
)

func AddShopper(c *gin.Context) {
	shopper := shopper.New()
	common.BindResponse(c,shopper)
	common.InsertDBWithContext(c,shopper)
	common.AfterService(c,nil,shopper)
}

func AddStockToShopper(c *gin.Context)  {
	var stock *shopper.ShopperStock = &shopper.ShopperStock{}
	common.BindResponse(c,stock)
	common.InsertDBWithContext(c,stock)
	common.AfterService(c,nil,stock)
}