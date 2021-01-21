package middleware

import (
	"context"
	"net/http"
)

func SessionFromCookie(r *http.Request) (token string, ok bool) {
	if tokencookie, err := r.Cookie("token"); err == nil && tokencookie.Value != "" {
		return tokencookie.Value, true
	}
	return "", false
}

type sessiontokenkey struct{}

func GetSessionToken(ctx context.Context) (string, bool) {
	if token, ok := ctx.Value(sessiontokenkey{}).(string); ok {
		return token, true
	}
	return "", false
}
func Session(n http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token, ok := SessionFromCookie(r)
		if ok {
			r = r.WithContext(context.WithValue(r.Context(), sessiontokenkey{}, token))
		}
		n.ServeHTTP(rw, r)
	})
}
