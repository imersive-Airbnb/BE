package routes

import (
	"airnbn/app/middlewares"
	_userData "airnbn/features/user/data"
	_userHandler "airnbn/features/user/handler"
	_userService "airnbn/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	// // Register middleware
	jwtMiddleware := middlewares.JWTMiddleware()

	// User Routes

	e.POST("/register", userHandlerAPI.RegisterUser)
	e.POST("/login", userHandlerAPI.Login)
	e.GET("/users", userHandlerAPI.CheckProfileByID, jwtMiddleware)
	e.PUT("/users/:id", userHandlerAPI.UpdateUserByID, jwtMiddleware)
}
