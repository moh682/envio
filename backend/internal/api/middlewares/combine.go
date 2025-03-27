package middlewares

import "net/http"

func Combine(nexts ...http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, next := range nexts {
			next.ServeHTTP(w, r)
		}
	})
}
