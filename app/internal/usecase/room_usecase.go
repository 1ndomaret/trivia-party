package usecase

import (
	"app/app/internal/domain"
	"app/app/internal/model/entity"
	"context"
)

type roomUsecase struct {
	roomRepo domain.RoomRepository
}

func NewRoomUsecase(roomRepo domain.RoomRepository) domain.RoomUsecase {
	return &roomUsecase{
		roomRepo: roomRepo,
	}
}

func (u *roomUsecase) Create(ctx context.Context, room *entity.Room) error {
	return nil
}

func Run(room *entity.Room) {

}
