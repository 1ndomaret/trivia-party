package usecase

import "app/app/internal/domain"

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return userUsecase{
		userRepo: userRepo,
	}
}
