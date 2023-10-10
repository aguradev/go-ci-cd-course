package controllers

import (
	"log"
	"net/http"
	"praktikum/middleware"
	"praktikum/models"
	"praktikum/repositories"
	"praktikum/resources"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authentication(c echo.Context) error {

	var UserRequest models.Users

	request := c.Bind(&UserRequest)

	if request != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, request.Error)
	}

	CheckEmailExists, ErrExp := repositories.CheckUserExistsRepositories(&UserRequest)

	if ErrExp != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrExp)
	}

	GenerateJWT, ErrJWT := middleware.GenerateToken(CheckEmailExists.Name, CheckEmailExists.Email)

	if ErrJWT != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, ErrJWT)
	}

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "Authentication Success",
		Data: map[string]interface{}{
			"name":  CheckEmailExists.Name,
			"email": CheckEmailExists.Email,
			"token": GenerateJWT,
		},
	})

}

// digunakan untuk menampilkan informasi user yang login melalui token
func TokenUserLoggedController(c echo.Context) error {

	token, valid := c.Get("user").(*jwt.Token)

	log.Println(token)

	if !valid {
		return echo.NewHTTPError(http.StatusBadRequest, "Token Missing Or Invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to claims token")
	}

	return c.JSON(http.StatusOK, resources.BaseResponse{
		Status:  http.StatusOK,
		Message: "User Logged",
		Data:    claims,
	})
}
