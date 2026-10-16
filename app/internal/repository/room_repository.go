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
