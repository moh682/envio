package middlewares

import "net/http"

type responseWriter struct {
	http.ResponseWriter
	wroteHeader bool
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	if !rw.wroteHeader {
		rw.wroteHeader = true
		rw.ResponseWriter.WriteHeader(statusCode)
	}
}

func Combine(nexts ...http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w}
		for _, next := range nexts {
			next.ServeHTTP(rw, r)
			if rw.wroteHeader {
				break
			}
		}
	})
}
