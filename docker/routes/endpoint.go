package routes

import (
	"docker/controllers"
	h "docker/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-jwt/v4"
)

func New() *echo.Echo{
	e := echo.New()
	h.LogMiddleware(e)
	//Login
	e.POST("/login", controllers.LoginUserController)
	
	//User
	JWT := e.Group("/auth")
	JWT.Use(echojwt.JWT([]byte("5347876725")))
	JWT.GET("/users", controllers.GetUsersController)
	JWT.GET("/users/:id", controllers.GetUserController)
	JWT.POST("/users", controllers.CreateUserController)
	JWT.DELETE("/users/:id", controllers.DeleteUserController)
	JWT.PUT("/users/:id", controllers.UpdateUserController)
	
	//Buku
	JWT.POST("/books", controllers.CreateBookController)
	JWT.GET("/books", controllers.GetBookController)
	JWT.GET("/books/:id", controllers.GetBookidController)
	JWT.PUT("/books/:id", controllers.UpdateBookController)
	JWT.DELETE("/books/:id", controllers.DeleteBookController)
	return e
}