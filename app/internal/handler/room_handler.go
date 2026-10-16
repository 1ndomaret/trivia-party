package handler

import (
	"app/app/internal/domain"
	"app/app/internal/helper"
	"app/app/internal/model/entity"

	"github.com/google/uuid"

	"github.com/labstack/echo/v5"
)

type RoomHandler struct {
	roomUsecase domain.RoomUsecase
}

func NewRoomHandler(roomUsecase domain.RoomUsecase) *RoomHandler {
	return &RoomHandler{
		roomUsecase: roomUsecase,
	}
}

func (h *RoomHandler) CreateRoom(c *echo.Context) error {
	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return helper.Unauthorized(c, "invalid or missing token")
	}

	room := entity.Room{
		UserID: userID,
	}

	if err := h.roomUsecase.Create(c.Request().Context(), &room); err != nil {
		return helper.InternalServerError(c, err.Error())
	}

	return helper.Created(c, room)
}

func (h *RoomHandler) UpdateRoom(c *echo.Context) error {

	return nil
}

func (h *RoomHandler) StartRoom(c *echo.Context) error {

	return nil
}

func (h *RoomHandler) JoinRoom(c *echo.Context) error {

	return nil
}
