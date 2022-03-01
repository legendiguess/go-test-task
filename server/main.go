package main

import (
	"context"
	"log"
	"net"
	"time"

	"test-task/cache"
	db "test-task/database"
	"test-task/domain"
	"test-task/logdb"
	pb "test-task/userserver"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServerServer
	database *db.Database
	cache    *cache.Cache
	logdb    *logdb.LogDB
	ctx      context.Context
}

func (s *server) AddUser(ctx context.Context, in *pb.User) (*pb.Response, error) {
	newUser := domain.User{Name: in.Name}

	s.database.AddUser(newUser)
	s.logdb.LogUser(newUser)

	return &pb.Response{Result: true}, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.User) (*pb.Response, error) {
	s.database.DeleteUser(domain.User{Name: in.Name})

	return &pb.Response{Result: true}, nil
}

func (s *server) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.Users, error) {
	users, err := s.cache.GetCacheUsers()
	if err != redis.Nil && err != nil {
		return &pb.Users{}, err
	}

	var names []string
	if err == redis.Nil {
		users = s.database.GetUsers()
		s.cache.CacheUsers(users)
	}

	for _, user := range users {
		names = append(names, user.Name)
	}

	return &pb.Users{Names: names}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	database, err := db.NewDatabse("host=localhost user=postgres password=postgres dbname=db port=5432 sslmode=disable TimeZone=Europe/Moscow")
	if err != nil {
		log.Fatalf("%v", err)
	}

	ctx := context.Background()

	cache := cache.NewCache(ctx, "localhost:6379", 60*time.Second)

	logdb := logdb.NewLogDB(ctx)

	server := server{database: database, cache: cache, logdb: logdb, ctx: ctx}

	pb.RegisterUserServerServer(s, &server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
