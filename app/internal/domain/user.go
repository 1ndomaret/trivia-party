package domain

import (
	"app/app/internal/model/dto"
	"app/app/internal/model/entity"
	"context"
	"errors"
)

var (
	ErrEmailExists = errors.New("email already registered")
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}

type UserUsecase interface {
	Register(ctx context.Context, req *dto.UserRegisterRequest) (*entity.User, error)
	Login(ctx context.Context, req *dto.UserLoginRequest) (string, error)
}
