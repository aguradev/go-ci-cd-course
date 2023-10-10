package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtCustomClaim struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(name, email string) (string, error) {

	var Payload jwtCustomClaim

	Payload.Name = name
	Payload.Email = email
	Payload.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 72))

	GeneratingToken := jwt.NewWithClaims(jwt.SigningMethodHS256, Payload)

	Token, TokenErr := GeneratingToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if TokenErr != nil {
		return "", TokenErr
	}

	return Token, nil

}

func AuthMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{
				"status":  http.StatusUnauthorized,
				"message": "missing or malformed jwt",
			})
		},
	})
}
