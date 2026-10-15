package handler

import (
	"app/app/internal/domain"
	"app/app/internal/helper"
	"app/app/internal/model/dto"
	"errors"

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

	return helper.Created(c, user)
}

func (h *UserHandler) Login(c *echo.Context) error {
	var req dto.UserLoginRequest

	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	token, err := h.userUsecase.Login(c.Request().Context(), &req)
	if err != nil {
		return helper.Unauthorized(c, "Invalid email or password")
	}

	return helper.Success(c, 200, "Login Success", dto.UserLoginResponse{
		Token: token,
	})
}
