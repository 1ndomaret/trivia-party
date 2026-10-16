package usecase

import (
	"app/app/internal/domain"
	"app/app/internal/helper"
	"app/app/internal/model/dto"
	"app/app/internal/model/entity"
	"context"
	"fmt"

	"github.com/google/uuid"
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
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	// TODO: CHECK FOR EXISTING ACTIVE ROOMS BEFORE CREATING A ROOM
	newCode, err := helper.Generate4CharRoomCode()
	if err != nil {
		return fmt.Errorf("Failed to generate room code: %s", err.Error())
	}

	room.RoomCode = newCode

	return u.roomRepo.Create(ctx, room)
}

func (u *roomUsecase) Update(ctx context.Context, userID uuid.UUID, roomCode string, req *dto.RoomUpdateRequest) (*entity.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	room, err := u.roomRepo.FindByRoomCode(ctx, roomCode)
	if err != nil {
		return nil, err
	}

	if room.UserID != userID {
		return nil, domain.ErrUnauthorized
	}

	if req.QuestionCategory != "" {
		room.QuestionCategory = req.QuestionCategory
	}
	if req.QuestionCount != 0 {
		room.QuestionCount = req.QuestionCount
	}
	if req.MaxPlayers != 0 {
		room.MaxPlayers = req.MaxPlayers
	}

	if room.QuestionCount < 1 || room.QuestionCount > 30 {
		return nil, domain.ErrInvalidQuestionCount
	}
	if room.MaxPlayers < 1 || room.MaxPlayers > 20 {
		return nil, domain.ErrInvalidMaxPlayers
	}

	if err := u.roomRepo.Update(ctx, room); err != nil {
		return nil, err
	}

	return room, nil
}

func (u *roomUsecase) Run(ctx context.Context, room *entity.Room) error {
	return nil
}
