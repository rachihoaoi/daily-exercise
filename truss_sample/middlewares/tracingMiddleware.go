package middlewares

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"net/http"
	"reflect"
)

func TracingMiddleWare(handler http.Handler) http.Handler {
	tracer := opentracing.GlobalTracer()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var span opentracing.Span
		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			span = tracer.StartSpan(r.RequestURI)
		} else {
			span = tracer.StartSpan(r.RequestURI, ext.RPCServerOption(spanCtx))
		}
		defer span.Finish()

		span.LogFields(reflectRequest(r)...)
		ctx := opentracing.ContextWithSpan(context.Background(), span)
		handler.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}

func reflectRequest(r *http.Request) []log.Field {
	fields := make([]log.Field, 0)
	vals := reflect.ValueOf(*r)
	types := reflect.TypeOf(*r)
	for i := 0; i < types.NumField(); i++ {
		switch vals.Field(i).Kind() {
		case reflect.String:
			fields = append(fields, log.String(types.Field(i).Name, vals.Field(i).String()))
		case reflect.Int64:
			fields = append(fields, log.Int64(types.Field(i).Name, vals.Field(i).Int()))
		}
	}
	return fields
}
