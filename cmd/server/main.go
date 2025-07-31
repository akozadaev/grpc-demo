package main

import (
	"context"
	"fmt"
	"github.com/akozadaev/grpc-demo/echo"
	helper "github.com/akozadaev/grpc-demo/pkg"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: "Echo: " + req.Message}, nil
}

func main() {
	network := helper.GetEnv("NETWORK")
	address := helper.GetEnv("PORT")

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
