package middleware

import "net/http"

func SessionFromCookie(r *http.Request) (token string, domain string, ok bool) {
	if tokencookie, err := r.Cookie("LWSSO_COOKIE_KEY"); err == nil {
		token = tokencookie.Value
	} else {
		return "", "", false
	}
	if domaincookie, err := r.Cookie("currentDomain"); err == nil {
		domain = domaincookie.Value
	} else {
		return "", "", false
	}
	return token, domain, true
}
