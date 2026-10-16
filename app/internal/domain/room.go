package domain

import (
	"app/app/internal/model/entity"
	"context"
)

type RoomRepository interface {
	Create(ctx context.Context, room *entity.Room) error
}

type RoomUsecase interface {
	Create(ctx context.Context, room *entity.Room) error
}
