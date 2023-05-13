package middlewares

import (
	"MiniProject_Rahman/project/constant"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CreateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(constant.SECRET_JWT)
	return token.SignedString(byteSecret)
}

func JWT() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &jwt.MapClaims{},
		SigningKey: []byte(constant.SECRET_JWT),
	}
	return middleware.JWTWithConfig(config)
}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		return userId
	}

	return 0
}
