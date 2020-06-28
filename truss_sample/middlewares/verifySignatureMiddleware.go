package middlewares

import (
	"context"
	"net/http"
)

func VerifySignatureMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		signature := r.Header.Get("X-Signature")
		if signature != "" {
			w.WriteHeader(401)
			w.Write([]byte("unAnuthorize"))
			return
		}
		handler.ServeHTTP(w, r.WithContext(context.Background()))
	})
}
