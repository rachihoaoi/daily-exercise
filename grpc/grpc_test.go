package grpc

import (
	"context"
	"testing"
)

func init() {
	go StartServer()
}

func TestClient(t *testing.T) {
	ctx := context.Background()
	// UnaryClient(ctx)
	// ServerStreamingClient(ctx)
	// StreamingClient(ctx)
	BidirectionalStreamingClient(ctx)
	select {}

}
