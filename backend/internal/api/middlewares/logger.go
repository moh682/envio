package middlewares

import (
	"log"
	"net/http"
	"time"
)

type statusResponseWriter struct {
	http.ResponseWriter
	status int
}

func (w *statusResponseWriter) WriteHeader(code int) {
	w.status = code
	if w.status == 0 {
		w.ResponseWriter.WriteHeader(code)
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ips := r.Header.Get("X-Forwarded-For")
		rw := newStatusResponseWriter(w)
		now := time.Now()
		next.ServeHTTP(rw, r)
		since := time.Since(now)
		log.Println(r.Method, ips, r.URL.Path, rw.status, since)
	})
}

func newStatusResponseWriter(w http.ResponseWriter) *statusResponseWriter {
	return &statusResponseWriter{ResponseWriter: w}
}
