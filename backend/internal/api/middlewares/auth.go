package middlewares

import (
	"net/http"

	"github.com/supertokens/supertokens-golang/recipe/session"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session.VerifySession(nil, next.ServeHTTP).ServeHTTP(w, r)
	})
}
