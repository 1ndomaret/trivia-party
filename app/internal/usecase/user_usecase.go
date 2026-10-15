package usecase

import (
	"app/app/internal/domain"
	"app/app/internal/model/dto"
	"app/app/internal/model/entity"
	"context"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

var timeOut = 2 * time.Second

func (u *userUsecase) Register(ctx context.Context, req *dto.UserRegisterRequest) (*entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeOut)
	defer cancel()

	existingUser, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	var user = entity.User{
		ID:       uuid.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := u.userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userUsecase) Login(ctx context.Context, req *dto.UserLoginRequest) (string, error) {
	return "", nil
}
