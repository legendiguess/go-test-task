package main

import (
	"context"
	"log"
	"time"

	pb "test-task/userserver"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err = c.AddUser(ctx, &pb.User{Name: "Arnold"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	_, err = c.AddUser(ctx, &pb.User{Name: "Aleksandr"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	_, err = c.AddUser(ctx, &pb.User{Name: "Platon"})
	if err != nil {
		log.Fatalf("%v", err)
	}
	_, err = c.DeleteUser(ctx, &pb.User{Name: "Platon"})
	if err != nil {
		log.Fatalf("%v", err)
	}

	u, err := c.GetUsers(ctx, &pb.GetUsersParams{})
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("users: %v", u)

}
