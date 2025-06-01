package repository

import (
	"context"
	"logger"
	"user-service/config"
	"user-service/model"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	CreateUser(ctx context.Context, user model.User) (err error)
}

type UserRepositoryImpl struct {
	Config *config.Config
	DB     *sqlx.DB
}

func ProvideUserRepository(cfg *config.Config, db *sqlx.DB) UserRepository {
	return &UserRepositoryImpl{
		Config: cfg,
		DB:     db,
	}
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	query := userQuery.selectUser + " WHERE email = $1"
	err = r.DB.GetContext(ctx, &user, query, email)
	if err != nil {
		logger.L().Error("[UserRepository][GetUserByEmail]",
			zap.Error(err),
			zap.String("email", email),
		)
		return
	}

	return user, err
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, user model.User) (err error) {
	query := userQuery.insertUser
	_, err = r.DB.ExecContext(ctx, query, user.Name, user.Email, user.PasswordHash)

	return err
}
