package router

import (
	"app/app/internal/handler"
	"app/app/internal/middleware"
	ws "app/app/internal/websocket"

	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

func Register(e *echo.Echo, userHandler *handler.UserHandler, roomHandler *handler.RoomHandler) {
	e.GET("/ws", ws.CreateSession)
	api := e.Group("/api/v1")

	public := api.Group("")

	public.POST("/users/register", userHandler.Register)
	public.POST("/users/login", userHandler.Login)

	private := api.Group("")
	private.Use(echojwt.WithConfig(middleware.JwtConfig()), middleware.ParseJwtClaims)

	private.POST("/room", roomHandler.CreateRoom)
	private.PUT("/room/:id", roomHandler.UpdateRoom)
	private.POST("/room/:id/start", roomHandler.StartRoom)

	private.POST("/room/:id/join", roomHandler.JoinRoom)
}
