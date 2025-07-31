package main

import (
	"context"
	"fmt"
	"github.com/akozadaev/grpc-demo/echo"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: "Echo: " + req.Message}, nil
}

func main() {
	envPath := "."
	envFileName := ".env"

	fullPath := envPath + "/" + envFileName

	if err := godotenv.Overload(fullPath); err != nil {
		log.Printf("[ERROR] failed with %+v", "No .env file found")
	}

	network := os.Getenv("NETWORK")
	address := os.Getenv("PORT")

	lis, err := net.Listen(network, fmt.Sprintf(":%s", address))
	if err != nil {
		log.Fatalf("не удалось запустить слушатель: %v", err)
	}

	grpcServer := grpc.NewServer()
	echo.RegisterEchoServiceServer(grpcServer, &server{})

	log.Println(fmt.Sprintf("gRPC сервер слушает на :%s", address))

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("ошибка запуска сервера: %v", err)
	}
}
