package tracing

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"io"

	jaegercfg "github.com/uber/jaeger-client-go/config"
	zapwraper "github.com/uber/jaeger-client-go/log/zap"
)

const (
	JAEGER_HOST = "127.0.0.1:6831"
)

type MDCarrier struct {
	metadata.MD
}

func NewJaegerTracer(serviceName string, logger *zap.Logger) (tracer opentracing.Tracer, closer io.Closer, err error) {
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: JAEGER_HOST,
		},
	}
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err = cfg.NewTracer(
		jaegercfg.Logger(zapwraper.NewLogger(logger)),
		jaegercfg.Metrics(jMetricsFactory))

	opentracing.SetGlobalTracer(tracer)
	if err != nil {
		grpclog.Errorf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	return
}

func (m MDCarrier) Set(key, val string) {
	// key = strings.ToLower(key)
	m.MD[key] = append(m.MD[key], val)
}

func (m MDCarrier) ForeachKey(handler func(key, val string) error) error {
	for k, strVal := range m.MD {
		for _, v := range strVal {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func ClientInterceptor(tracer opentracing.Tracer) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		// get parent span from context
		var parentCtx opentracing.SpanContext
		parentSpan := opentracing.SpanFromContext(ctx)
		if parentSpan != nil {
			parentCtx = parentSpan.Context()
		}
		span := tracer.StartSpan(
			method,
			opentracing.ChildOf(parentCtx),
			opentracing.Tag{string(ext.Component), "gRPC Client"},
			ext.SpanKindRPCClient,
		)
		defer span.Finish()
		// get md from ctx
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}
		if err = tracer.Inject(span.Context(), opentracing.TextMap, MDCarrier{md}); err != nil {
			return err
		}
		if err = invoker(metadata.NewOutgoingContext(ctx, md), method, req, reply, cc, opts...); err != nil {
			return err
		}
		return err
	}
}

func ServerInterceptor(tracer opentracing.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		}
		spanContext, err := tracer.Extract(
			opentracing.TextMap,
			MDCarrier{md},
		)

		if err != nil && err != opentracing.ErrSpanContextNotFound {
			grpclog.Errorf("extract from metadata err: %v", err)
		} else {
			span := tracer.StartSpan(
				info.FullMethod,
				ext.RPCServerOption(spanContext),
				opentracing.Tag{Key: string(ext.Component), Value: "gRPC Server"},
				ext.SpanKindRPCServer,
			)
			defer span.Finish()

			ctx = opentracing.ContextWithSpan(ctx, span)
		}
		return handler(ctx, req)
	}
}
