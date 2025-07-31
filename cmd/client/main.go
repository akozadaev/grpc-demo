package main

import (
	"context"
	"fmt"
	"github.com/akozadaev/grpc-demo/echo"
	helper "github.com/akozadaev/grpc-demo/pkg"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	address := helper.GetEnv("PORT")
	target := fmt.Sprintf("localhost:%s", address)

	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}

	defer conn.Close()

	c := echo.NewEchoServiceClient(conn)

	msg := "Hello, gRPC!"
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Echo(ctx, &echo.EchoRequest{Message: msg})
	if err != nil {
		log.Fatalf("ошибка вызова Echo: %v", err)
	}
	log.Printf("Ответ: %s", r.Message)
}
