package cmd

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"

	"github.com/rachihoaoi/daily-exercise.git/tracing"
	helloHd "github.com/rachihoaoi/daily-exercise.git/truss_sample/internal/hello/handlers"
	helloSvc "github.com/rachihoaoi/daily-exercise.git/truss_sample/internal/hello/svc"
	"github.com/rachihoaoi/daily-exercise.git/truss_sample/middlewares"
	"github.com/rachihoaoi/daily-exercise.git/truss_sample/pbs"
)

func Run() {
	// init zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Errorf("fail to new zap logger with error: %v", err))
	}
	defer logger.Sync()

	// init tracer
	tracer, closer, err := tracing.NewJaegerTracer("main", logger)
	if err != nil {
		panic(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	endpoints := NewHelloEndpoint(tracer, logger, nil, nil)

	go MountGRPCServer(func(server *grpc.Server) {
		pb.RegisterGreetingServer(server, helloSvc.MakeGRPCServer(endpoints))
	})

	go MountHttpServer(
		helloSvc.MakeHTTPHandler(endpoints).(*mux.Router),
	)

	select {}

}

func NewHelloEndpoint(tracer opentracing.Tracer, logger *zap.Logger, mysqlClient *sql.DB, redisClient *redis.Client) helloSvc.Endpoints {
	var service pb.GreetingServer

	service = helloHd.NewService(tracer, logger, mysqlClient, redisClient)
	service = helloHd.WrapService(service)

	var (
		sayHelloEndpoint = helloSvc.MakeSayHelloEndpoint(service)
	)

	endpoints := helloSvc.Endpoints{
		SayHelloEndpoint: sayHelloEndpoint,
	}

	endpoints = helloHd.WrapEndpoints(tracer, logger, endpoints)
	return endpoints
}

func MountGRPCServer(registerFunc func(server *grpc.Server)) {
	var blockCh = make(chan bool)

	server := grpc.NewServer(grpc.UnaryInterceptor(tracing.ServerInterceptor(opentracing.GlobalTracer())))
	registerFunc(server)

	listener, err := net.Listen("tcp", ":1111")
	if err != nil {
		panic(err)
	}

	if err = server.Serve(listener); err != nil {
		panic(err)
	}
	blockCh <- true
}

func MountHttpServer(subRoutes ...*mux.Router) *mux.Router {
	var blockCh = make(chan error)
	rootRouter := mux.NewRouter().PathPrefix("").Subrouter()
	for _, subRouter := range subRoutes {
		subRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) (err error) {
			methods, _ := route.GetMethods()
			if methods[0] == "" {
				return
			}
			path, _ := route.GetPathTemplate()
			rootRouter.
				Methods(methods...).Path(path).
				Handler(route.GetHandler())
			return
		})
	}
	rootRouter.Use(middlewares.TracingMiddleWare)
	rootRouter.Use(middlewares.VerifySignatureMiddleware)

	blockCh <- http.ListenAndServe("127.0.0.1:1234", rootRouter)
	return rootRouter
}
