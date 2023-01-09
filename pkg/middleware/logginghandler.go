package middleware

import (
	"net/http"
	"strconv"

	"github.com/vladisvrau/FamilyTree/lib/log"
)

func LoggingHandler(next http.HandlerFunc, logger log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := loggingResponseWriter{w, 0}

		next.ServeHTTP(&writer, r)
		logger.Info(strconv.Itoa(writer.statusCode), r.Method, "\t", r.URL.Path)
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (l *loggingResponseWriter) WriteHeader(status int) {
	l.ResponseWriter.WriteHeader(status)
	l.statusCode = status
}
