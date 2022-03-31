package routes

import (
	_authHandler "be7/layered/delivery/handler/auth"
	_bookHandler "be7/layered/delivery/handler/book"
	_userHandler "be7/layered/delivery/handler/user"
	_middlewares "be7/layered/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, bh *_bookHandler.BookHandler) {
	// user
	e.GET("/users", uh.GetAllHandler(), _middlewares.JWTMiddleware())
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdatedUserHandler(), _middlewares.JWTMiddleware())

	// book
	e.GET("/books", bh.GetAllHandler())
	e.POST("/books", bh.CreateBookHandler(), _middlewares.JWTMiddleware())
	e.GET("/books/:id", bh.GetBookHandler())
	e.DELETE("/books/:id", bh.DeleteBookHandler(), _middlewares.JWTMiddleware())
	e.PUT("/books/:id", bh.UpdatedBookHandler(), _middlewares.JWTMiddleware())

}

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}
