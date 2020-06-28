// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: c575929542
// Version Date: 2020年 6月17日 星期三 04时40分15秒 UTC

package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/rachihoaoi/daily-exercise.git/truss_sample/pbs"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC GreetingServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.GreetingServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// greeting

		sayhello: grpctransport.NewServer(
			endpoints.SayHelloEndpoint,
			DecodeGRPCSayHelloRequest,
			EncodeGRPCSayHelloResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the GreetingServer interface
type grpcServer struct {
	sayhello grpctransport.Handler
}

// Methods for grpcServer to implement GreetingServer interface

func (s *grpcServer) SayHello(ctx context.Context, req *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	_, rep, err := s.sayhello.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SayHelloResponse), nil
}

// Server Decode

// DecodeGRPCSayHelloRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sayhello request to a user-domain sayhello request. Primarily useful in a server.
func DecodeGRPCSayHelloRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SayHelloRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCSayHelloResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sayhello response to a gRPC sayhello reply. Primarily useful in a server.
func EncodeGRPCSayHelloResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.SayHelloResponse)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
