package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"net"
	"time"

	"github.com/rachihoaoi/daily-exercise.git/grpc/pbs"
)

type Server struct{}

func (s Server) UnaryService(ctx context.Context, req *pbs.Req) (resp *pbs.Resp, err error) {
	return &pbs.Resp{
		Result: "get: " + req.GetMessage(),
	}, nil
}

func (s Server) ServerSideStreaming(req *pbs.Req, stream pbs.Sample_ServerSideStreamingServer) (err error) {
	for {
		time.Sleep(time.Second)
		_ = stream.Send(&pbs.Resp{
			Result: req.Message + time.Now().String(),
		})
		if time.Now().Second() == 33 {
			break
		}
	}
	return nil
}

func (s Server) ClientSideStreaming(stream pbs.Sample_ClientSideStreamingServer) (err error) {
	for {
		req, err := stream.Recv()
		if err != nil {
			return stream.SendAndClose(&pbs.Resp{Result: "end: " + err.Error()})
		}
		fmt.Println(req.Message)
	}
}

func (s Server) BidirectionalStreaming(stream pbs.Sample_BidirectionalStreamingServer) error {
	for {
		stream.Send(&pbs.Resp{Result: "get"})
		r, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		fmt.Println(r.Message)
	}
}

func StartServer() {
	var blockCh = make(chan bool)
	server := grpc.NewServer()
	pbs.RegisterSampleServer(server, &Server{})

	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		panic(err)
	}

	if err = server.Serve(listener); err != nil {
		panic(err)
	}

	blockCh <- true
}
