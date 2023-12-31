package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// LOG middleware
// ? berfungsi untuk mengecek sebuah method yang sedang berjalan. response data merupakan status dan milisecond dari data tersebut.

func LogMiddleware(e *echo.Echo) {

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human} \n",
	}))

}
