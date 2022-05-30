package Routes

import (
	"user-service/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/", Controllers.Register)

	router.GET("/")

	router.GET("/verification/code/:whatsapp_no", Controllers.RequestVerificationCode)

	router.POST("/verification/code", Controllers.VerifyVerificationCode)

	router.POST("/verification/password", Controllers.VerifyPassword)

	router.PATCH("/update/name", Controllers.UpdateName)

	router.PATCH("/update/password", Controllers.UpdatePassword)

	router.POST("/form")

	return router
}
