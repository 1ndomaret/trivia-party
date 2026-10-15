package handler

import (
	"app/app/internal/domain"
	"app/app/internal/helper"
	"app/app/internal/model/dto"
	"errors"
	"net/http"

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
	var req dto.UserRegisterRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	if err := c.Validate(&req); err != nil {
		return helper.BadRequest(c, helper.FormatValidationError(err))
	}

	user, err := h.userUsecase.Register(c.Request().Context(), &req)
	if err != nil {
		if errors.Is(err, domain.ErrEmailExists) {
			return helper.BadRequest(c, "Email already registered")
		}

		return helper.InternalServerError(c, err.Error())
	}

	return helper.Success(c, http.StatusCreated, "Succesfully Registered", user)
}
