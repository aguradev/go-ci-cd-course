package routes

import (
	"net/http"
	"praktikum/controllers"
	m "praktikum/middleware"

	"github.com/labstack/echo/v4"
)

func Route() {

	e := echo.New()

	// check logging in endpoint
	m.LogMiddleware(e)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Selamat datang di APIku",
		})
	})

	e.POST("/auth/login", controllers.Authentication)
	e.POST("/users", controllers.CreateUserController)

	JWTauth := e.Group("", m.AuthMiddleware())

	JWTauth.GET("/users", controllers.GetAllUsersController)
	JWTauth.GET("/users/token-logged", controllers.TokenUserLoggedController)

	JWTauth.GET("/users/:id", controllers.GetUserByIdController)
	JWTauth.DELETE("/users/:id", controllers.DeleteUserController)
	JWTauth.PUT("/users/:id", controllers.UpdateUserController)

	JWTauth.GET("/books", controllers.GetBooksController)
	JWTauth.GET("/books/:id", controllers.GetBookByIdController)
	JWTauth.POST("/books", controllers.CreateBookController)
	JWTauth.DELETE("/books/:id", controllers.DeleteBookController)
	JWTauth.PUT("/books/:id", controllers.UpdateBookController)

	JWTauth.GET("/blogs", controllers.GetAllBlogController)
	JWTauth.GET("/blogs/:id", controllers.GetBlogByIdController)
	JWTauth.POST("/blogs", controllers.CreateBlogController)
	JWTauth.DELETE("/blogs/:id", controllers.DeleteBlogController)
	JWTauth.PUT("/blogs/:id", controllers.UpdateBlogController)

	e.Logger.Fatal(e.Start(":1234"))

}
