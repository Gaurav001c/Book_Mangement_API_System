package routes

import (
	authication "github.com/GAURAV/BookApiTask/authentication"
	"github.com/GAURAV/BookApiTask/controller"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.POST("/register", controller.Register)
	e.POST("/login", controller.Login)

	e.GET("/categories", controller.GetCategories)
	e.POST("/categories", controller.CreateCategory, authication.JWTMiddleware())

	e.GET("/books", controller.GetBooks)
	e.POST("/books", controller.CreateBook, authication.JWTMiddleware())
	e.PUT("/books/:id", controller.UpdateBook, authication.JWTMiddleware())
	e.DELETE("/books/:id", controller.DeleteBook, authication.JWTMiddleware())
}
