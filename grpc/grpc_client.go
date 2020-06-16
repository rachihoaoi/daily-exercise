package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"strconv"
	"time"

	"github.com/rachihoaoi/daily-exercise.git/grpc/pbs"
)

func GetGRPCConnection(target string, opts []grpc.DialOption) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(target, opts...)
	if err != nil {
		panic(err)
	}
	return conn, err
}

func UnaryClient(ctx context.Context) {
	conn, _ := GetGRPCConnection(":1111", []grpc.DialOption{grpc.WithInsecure()})
	defer conn.Close()

	client := pbs.NewSampleClient(conn)

	resp, _ := client.UnaryService(ctx, &pbs.Req{Message: "test UnaryClient"})
	fmt.Println(resp)
}

func ServerStreamingClient(ctx context.Context) {
	conn, _ := GetGRPCConnection(":1111", []grpc.DialOption{grpc.WithInsecure()})
	defer conn.Close()

	client := pbs.NewSampleClient(conn)

	stream, _ := client.ServerSideStreaming(context.Background(), &pbs.Req{
		Message: "message",
	})
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println(resp)
	}
}

func StreamingClient(ctx context.Context) {
	conn, _ := GetGRPCConnection(":1111", []grpc.DialOption{grpc.WithInsecure()})
	defer conn.Close()

	client := pbs.NewSampleClient(conn)

	stream, _ := client.ClientSideStreaming(ctx)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		stream.Send(&pbs.Req{Message: strconv.Itoa(i)})
	}
	fmt.Println(stream.CloseAndRecv())
}

func BidirectionalStreamingClient(ctx context.Context) {
	conn, _ := GetGRPCConnection(":1111", []grpc.DialOption{grpc.WithInsecure()})
	defer conn.Close()

	client := pbs.NewSampleClient(conn)
	stream, _ := client.BidirectionalStreaming(ctx)
	for i := 0; i < 10; i++ {
		stream.Send(&pbs.Req{Message: strconv.Itoa(i)})
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(resp)
	}
	stream.CloseSend()
}
