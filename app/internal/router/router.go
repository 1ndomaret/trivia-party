package router

import (
	"app/app/internal/handler"
	"app/app/internal/socket"

	"github.com/labstack/echo/v5"
)

func Register(e *echo.Echo, userHandler *handler.UserHandler) {
	e.GET("/ws", socket.CreateSession)
	api := e.Group("/api/v1")

	public := api.Group("")
	public.POST("/users/register", userHandler.Register)
	public.POST("/users/login", userHandler.Login)
}
