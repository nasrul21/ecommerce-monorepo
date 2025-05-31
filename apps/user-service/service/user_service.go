package service

import (
	"context"
	"errors"
	"user-service/config"
	"user-service/model"
	"user-service/proto/userpb"
	"user-service/repository"
	"user-service/util"

	"github.com/google/uuid"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (res *userpb.RegisterUserResponse, err error)
}

type userService struct {
	cfg      *config.Config
	userRepo repository.UserRepository
}

func ProvideUserService(cfg *config.Config, userRepo repository.UserRepository) UserService {
	return &userService{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (s *userService) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (res *userpb.RegisterUserResponse, err error) {
	existingUser, _ := s.userRepo.GetUserByEmail(ctx, req.GetEmail())
	if existingUser.ID != uuid.Nil {
		err = errors.New("email already registered")
		return
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return
	}

	user := model.User{
		Name:         req.GetName(),
		Email:        req.GetEmail(),
		PasswordHash: hashedPassword,
	}

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		err = errors.New("failed to create user")
		return
	}

	res = &userpb.RegisterUserResponse{
		Ok: true,
	}
	return
}
