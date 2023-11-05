package main

import (
	pb "client/pb/client/proto"
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
		return
	}
	client := pb.NewUserServiceClient(conn)

	created, err := client.Create(context.Background(), &pb.User{
		Name:    "Sample User",
		Email:   "sample@gmail.com",
		Address: proto.String("21 Symphony Apartment"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created User with Id - {%d}\n", created.Id)

	user, err := client.Get(context.Background(), &pb.UserId{Id: created.Id})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found User - {%s}\n", user)
}
