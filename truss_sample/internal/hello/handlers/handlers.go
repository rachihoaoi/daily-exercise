package handlers

import (
	"context"
	"database/sql"
	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/rachihoaoi/daily-exercise.git/truss_sample/pbs"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService(tracer opentracing.Tracer, logger *zap.Logger, mysqlDB *sql.DB, redisClient *redis.Client) pb.GreetingServer {
	return greetingService{tracer, logger, mysqlDB, redisClient}
}

type greetingService struct {
	tracer      opentracing.Tracer
	logger      *zap.Logger
	mysqlDB     *sql.DB
	redisClient *redis.Client
}

// SayHello implements Service.
func (s greetingService) SayHello(ctx context.Context, in *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	var resp pb.SayHelloResponse
	resp = pb.SayHelloResponse{
		Sentence: "wocao",
	}
	return &resp, nil
}
