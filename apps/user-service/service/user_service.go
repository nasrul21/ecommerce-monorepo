package service

import (
	"context"
	"database/sql"
	"errors"
	"user-service/config"
	"user-service/model"
	"user-service/proto/userpb"
	"user-service/repository"
	"user-service/util"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserService interface {
	RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (res *userpb.RegisterUserResponse, err error)
	LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (res *userpb.LoginUserResponse, err error)
}

type UserServiceImpl struct {
	cfg      *config.Config
	userRepo repository.UserRepository
}

func ProvideUserService(cfg *config.Config, userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		cfg:      cfg,
		userRepo: userRepo,
	}
}

func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (res *userpb.RegisterUserResponse, err error) {
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

func (s *UserServiceImpl) LoginUser(ctx context.Context, req *userpb.LoginUserRequest) (res *userpb.LoginUserResponse, err error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.GetEmail())
	if err != nil && err != sql.ErrNoRows {
		zap.L().Error("[UserRepository][LoginUser] failed to retreive user data",
			zap.Error(err),
			zap.Any("user.email", req.GetEmail()),
		)
		return
	}

	err = util.ComparePassword(user.PasswordHash, req.GetPassword())
	if err != nil {
		err = errors.New("invalid password")
		return
	}

	token, err := util.GenerateJWT(
		user.ID.String(),
		user.Email,
		s.cfg.Auth.Token.Expired,
		s.cfg.Auth.Token.Secret,
	)
	if err != nil {
		zap.L().Error("[UserService][LoginUser] failed generate token",
			zap.Error(err),
			zap.String("user.id", user.ID.String()),
		)
		err = errors.New("failed to generate token")
		return
	}

	res = &userpb.LoginUserResponse{
		Token: token,
	}

	return
}
