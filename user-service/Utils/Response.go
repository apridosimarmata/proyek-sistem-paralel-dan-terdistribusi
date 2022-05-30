package Utils

import (
	"user-service/Models"

	"github.com/gin-gonic/gin"
)

func MakeResponse(c *gin.Context, statusCode int, message *string, result any) {
	c.Header("Content-Type", "application/json")

	c.JSON(
		200,
		Models.Response{
			Meta: Models.Meta{
				Code:    statusCode,
				Message: message,
			},
			Result: result,
		},
	)
}
