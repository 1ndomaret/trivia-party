package repository

import (
	"app/app/internal/domain"
	"app/app/internal/model/entity"
	"context"

	"gorm.io/gorm"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) domain.RoomRepository {
	return &roomRepository{
		db: db,
	}
}

func (r *roomRepository) Create(ctx context.Context, room *entity.Room) error {
	return r.db.WithContext(ctx).Create(room).Error
}

func (r *roomRepository) FindByRoomCode(ctx context.Context, roomCode string) (*entity.Room, error) {
	var room entity.Room

	if err := r.db.WithContext(ctx).Where("room_code = ? and status != ?", roomCode, entity.RoomFinished).First(&room).Error; err != nil {
		return nil, err
	}

	return &room, nil
}

func (r *roomRepository) Update(ctx context.Context, room *entity.Room) error {
	return r.db.WithContext(ctx).Save(room).Error
}
