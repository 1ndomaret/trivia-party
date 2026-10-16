package domain

import (
	"app/app/internal/model/dto"
	"app/app/internal/model/entity"
	"context"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized         = errors.New("forbidden")
	ErrInvalidQuestionCount = errors.New("question count must be between 1 and 30")
	ErrInvalidMaxPlayers    = errors.New("max players must be between 1 and 20")
)

type RoomRepository interface {
	Create(ctx context.Context, room *entity.Room) error
	FindByRoomCode(ctx context.Context, roomCode string) (*entity.Room, error)
	Update(ctx context.Context, room *entity.Room) error
}

type RoomUsecase interface {
	Create(ctx context.Context, room *entity.Room) error
	Update(ctx context.Context, userID uuid.UUID, roomCode string, req *dto.RoomUpdateRequest) (*entity.Room, error)
	Run(ctx context.Context, room *entity.Room) error
}
