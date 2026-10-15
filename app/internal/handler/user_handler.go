package handler

import (
	"app/app/internal/domain"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) Register(c *echo.Context) error {
	return nil
}
