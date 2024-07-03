package middlewares

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(rw http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(rw, r.ProtoMajor)
			start := time.Now()

			reqID := r.Context().Value(middleware.RequestIDKey)

			reqIDStr, ok := reqID.(string)
			if !ok {
				logger.Info("request started", "request-id", "unknown")
			}

			// "bytes", ww.BytesWritten(),

			defer func() {
				logger.Info("request completed",
					"request-id", reqIDStr,
					"status", ww.Status(),
					"method", r.Method,
					"path", r.URL.Path,
					"query", r.URL.RawQuery,
					"ip", r.RemoteAddr,
					"user-agent", r.UserAgent(),
					"latency-ms", time.Since(start),
				)
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
