package handler

import (
	"app/app/internal/domain"
	"app/app/internal/helper"
	"app/app/internal/model/dto"
	"app/app/internal/model/entity"
	"net/http"

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
	userID, ok := c.Get("userID").(uuid.UUID)
	if !ok {
		return helper.Unauthorized(c, "invalid or missing token")
	}

	roomCode := c.Param("code")
	if roomCode == "" {
		return helper.BadRequest(c, "Missing room code")
	}

	var req dto.RoomUpdateRequest
	if err := c.Bind(&req); err != nil {
		return helper.BadRequest(c, err.Error())
	}

	room, err := h.roomUsecase.Update(c.Request().Context(), userID, roomCode, &req)
	if err != nil {
		switch err {
		case domain.ErrUnauthorized:
			return helper.Unauthorized(c, "You do not have permission")
		case domain.ErrInvalidMaxPlayers:
			return helper.BadRequest(c, "Player count must be between or equal to 1 and 20")
		case domain.ErrInvalidQuestionCount:
			return helper.BadRequest(c, "Question count must be between or equal to 1 and 30")
		default:
			return helper.InternalServerError(c, err.Error())
		}
	}

	return helper.Success(c, http.StatusOK, "Room Updated", room)
}

func (h *RoomHandler) StartRoom(c *echo.Context) error {

	return nil
}

func (h *RoomHandler) JoinRoom(c *echo.Context) error {

	return nil
}
