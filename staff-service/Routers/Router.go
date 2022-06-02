package routers

import (
	controllers "staff-service/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// r is a router
	r := gin.Default()

	r.GET("/:staffUID", controllers.GetStaffByUID)
	r.POST("/", controllers.AuthenticateStaff)
	r.GET("/authorize", controllers.ValidateToken)
	return r
}
