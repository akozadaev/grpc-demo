package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/akozadaev/grpc-demo/echo"
	helper "github.com/akozadaev/grpc-demo/pkg"
	"google.golang.org/grpc"
)

type server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{Message: "Echo: " + req.Message}, nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	network := helper.GetEnv("NETWORK")
	address := helper.GetEnv("PORT")

	grpcServer := grpc.NewServer()
	echo.RegisterEchoServiceServer(grpcServer, &server{})

	go func() {
		lis, err := net.Listen(network, fmt.Sprintf(":%s", address))
		if err != nil {
			log.Fatalf("не удалось запустить слушатель: %v", err)
		}

		log.Println(fmt.Sprintf("gRPC сервер слушает на :%s", address))

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("ошибка запуска сервера: %v", err)
		}
	}()

	// Ждем сигнала завершения
	<-ctx.Done()

	// Завершаем работу сервера с таймаутом
	fmt.Println("Начинается graceful shutdown...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Graceful shutdown gRPC сервера
	done := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("gRPC сервер успешно завершен")
	case <-shutdownCtx.Done():
		fmt.Println("Таймаут graceful shutdown, принудительное завершение")
		grpcServer.Stop()
	}
}
