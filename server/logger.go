package server

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (w *ResponseRecorder) WriteHeader(code int) {
	w.StatusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *ResponseRecorder) Write(body []byte) (int, error) {
	w.Body = body
	return w.ResponseWriter.Write(body)
}

func (server *Server) logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()

		writer := &ResponseRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		next.ServeHTTP(writer, r)
		duration := time.Since(t1)

		logger := log.Info()
		if writer.StatusCode != http.StatusOK {
			logger = log.Error().Bytes("body", writer.Body)
		}

		logger.Str("method", r.Method).
			Str("uri", r.RequestURI).
			Int("status_code", writer.StatusCode).
			Str("status_text", http.StatusText(writer.StatusCode)).
			Dur("duration", duration).
			Msg("HTTP request")

	})
}
