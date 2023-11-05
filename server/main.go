package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"server/config"
	pb "server/pb/server/proto"
)

var DB *gorm.DB

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// Create DB Connection
	DB, _ = config.ConnectSql()

	// Create a gRPC Server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Create(ctx context.Context, user *pb.User) (*pb.User, error) {
	userTable := DB.Table("users")
	userTable.Create(user)
	return user, nil
}

func (s *server) Get(ctx context.Context, userId *pb.UserId) (*pb.User, error) {
	user := pb.User{}
	tx := DB.Table("users").Find(&user, userId.Id)
	if tx.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

type server struct {
	pb.UserServiceServer
}
