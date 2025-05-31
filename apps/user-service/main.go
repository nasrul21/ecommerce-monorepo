package main

import (
	"fmt"
	"log"
	"net"
	"user-service/config"
	"user-service/db"
	userGrpcHandler "user-service/handler/grpc"
	"user-service/proto/userpb"
	"user-service/repository"
	"user-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate protoc --go_out=. --go-grpc_out=. -I=../../packages/protos/user ../../packages/protos/user/user_payload.proto
func main() {
	cfg := config.LoadConfig()

	postgres := db.NewPostgres(cfg)

	userRepo := repository.ProvideUserRepository(cfg, postgres)
	userService := service.ProvideUserService(cfg, userRepo)
	userHandler := userGrpcHandler.ProvideUserHandler(userService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userHandler)

	reflection.Register(grpcServer)

	log.Printf("gRPC server running on :%d\n", cfg.GRPC.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
