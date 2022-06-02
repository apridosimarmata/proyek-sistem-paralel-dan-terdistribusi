package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GenerateJWT(username string, tokenType int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username

	if tokenType == TOKEN {
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	} else {
		claims["exp"] = time.Now().Add(time.Minute * 43800).Unix()
	}

	tokenString, err := token.SignedString([]byte(GoDotEnvVariable("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(tokenString string, username *string) int {

	var code = http.StatusOK

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			code = http.StatusNotAcceptable
		}
		checkExpired := token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), false)
		if !checkExpired {
			code = http.StatusGone
		}

		return []byte(GoDotEnvVariable("SECRET_KEY")), nil
	})

	if err != nil {
		if code != http.StatusGone {
			code = http.StatusBadRequest
		}
		fmt.Println(err.Error())
	}

	*username, _ = token.Claims.(jwt.MapClaims)["username"].(string)

	if token.Valid {
		return code
	}

	return code
}
