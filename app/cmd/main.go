package main

import (
	"app/app/internal/config"
	"app/app/internal/handler"
	"app/app/internal/helper"
	"app/app/internal/model/entity"
	"app/app/internal/repository"
	"app/app/internal/router"
	"app/app/internal/usecase"
	"log"

	"github.com/labstack/echo/v5"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	roomRepository := repository.NewRoomRepository(db)
	roomUsecase := usecase.NewRoomUsecase(roomRepository)
	roomHandler := handler.NewRoomHandler(roomUsecase)

	e := echo.New()
	e.Validator = helper.NewCustomValidator()

	router.Register(e,
		userHandler,
		roomHandler,
	)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
