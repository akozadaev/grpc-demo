package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akozadaev/grpc-demo/echo"
	helper "github.com/akozadaev/grpc-demo/pkg"

	"google.golang.org/grpc"
)

func main() {
	address := helper.GetEnv("PORT")
	target := fmt.Sprintf("localhost:%s", address)

	conn, err := grpc.NewClient(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("не удалось подключиться: %v", err)
	}

	defer conn.Close()

	c := echo.NewEchoServiceClient(conn)
	cm := echo.NewMathServiceClient(conn)

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

	a, err := cm.Add(ctx, &echo.NumbersRequest{A: 3, B: 1})
	log.Printf("Сумма: %s", a)

	d, err := cm.Divide(ctx, &echo.NumbersRequest{A: 3, B: 1})
	if err != nil {
		log.Fatalf("ошибка: %v", err)
	}
	log.Printf("Деление: %s", d)

	m, err := cm.Multiply(ctx, &echo.NumbersRequest{A: 3, B: 1})
	if err != nil {
		log.Fatalf("ошибка: %v", err)
	}
	log.Printf("Умножение: %s", m)

	s, err := cm.Subtract(ctx, &echo.NumbersRequest{A: 3, B: 1})
	if err != nil {
		log.Fatalf("ошибка: %v", err)
	}
	log.Printf("Вычитание: %s", s)

}
