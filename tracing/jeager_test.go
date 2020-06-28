package tracing

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"

	pb "github.com/rachihoaoi/daily-exercise.git/truss_sample/pbs"
)

func TestJaeger(t *testing.T) {
	// init zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("fail to new zap logger with error: %v", err))
	}
	defer logger.Sync()

	tracer, closer, err := NewJaegerTracer("client", logger)
	defer closer.Close()
	if err != nil {
		log.Printf("NewJaegerTracer err: %v", err.Error())
	}
	// Set up a connection to the server.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, "127.0.0.1:1111", grpc.WithInsecure(), grpc.WithUnaryInterceptor(ClientInterceptor(tracer)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetingClient(conn)

	resp, err := client.SayHello(ctx, &pb.SayHelloRequest{PersonName: "aaa"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
