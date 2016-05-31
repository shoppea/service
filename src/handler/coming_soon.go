package handler

import (
	"github.com/gin-gonic/gin"
	"coming_soon"
	"common"
	"throw"
)


// Handler to accept user
// Email from coming soon
// page and will add it to
// database
func ComingSoon(c *gin.Context) {
	var u coming_soon.User
	common.BindResponse(c,&u)
	err := u.RegisterForNotify()
	if err != nil {
		throw.ErrorDB(c,err)
	}
	throw.SuccessCreated(c,u)
}
