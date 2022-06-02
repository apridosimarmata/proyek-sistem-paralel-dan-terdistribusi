package controllers

import (
	"net/http"

	models "staff-service/Models"
	utils "staff-service/Utils"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func GetStaffByUID(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func ValidateToken(c *gin.Context) {
	var username string
	var token string

	cookie, err := c.Request.Cookie("token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
	}

	token = cookie.Value

	var code = utils.ValidateJWTToken(token, &username)

	var response models.BaseResponse

	switch code {
	case 200:
		response.Message = "success"
		response.Result = username
	case 400:
		response.Message = "bad request"
	case 410:
		response.Message = "token is expired"
	default:
		response.Message = "not acceptable"
	}

	c.JSON(code, response)
}

func AuthenticateStaff(c *gin.Context) {
	var staff models.Staff
	var login models.Login
	var response models.BaseResponse

	err := c.BindJSON(&login)

	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = models.GetStaffByUsername(&staff, login.Username)

	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(login.Password)) != nil {
		response.Message = "wrong credentials"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	var token models.JWT

	err = token.Create(staff.Username)

	if err != nil {
		response.Message = err.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Message = "success"

	response.Result = token

	c.JSON(http.StatusOK, response)
}
