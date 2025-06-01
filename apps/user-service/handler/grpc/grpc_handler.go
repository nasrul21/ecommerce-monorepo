package grpc

import (
	"context"
	"user-service/proto/userpb"
	"user-service/service"
)

type userHandler struct {
	userpb.UnimplementedUserServiceServer
	svc service.UserService
}

func ProvideUserHandler(svc service.UserService) userpb.UserServiceServer {
	return &userHandler{
		svc: svc,
	}
}

func (h *userHandler) Register(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	res, err := h.svc.RegisterUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *userHandler) Login(ctx context.Context, req *userpb.LoginUserRequest) (*userpb.LoginUserResponse, error) {
	res, err := h.svc.LoginUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
