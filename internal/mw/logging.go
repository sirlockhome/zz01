package mw

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func RequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestID := fmt.Sprintf("%d", time.Now().UnixNano())

		ctx := log.With().Str("request_id", requestID).Logger().WithContext(r.Context())

		wrapped := wrapResponseWriter(w)

		next.ServeHTTP(wrapped, r.WithContext(ctx))

		dur := time.Since(start)
		log.Info().
			Str("request_id", requestID).
			Str("method", r.Method).
			Str("path", r.URL.EscapedPath()).
			Int("status_code", wrapped.status).
			Dur("duration", dur).
			Msg("complate request")
	})
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
