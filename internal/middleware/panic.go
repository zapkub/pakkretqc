package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

type doneWriter struct {
	http.ResponseWriter
	isheaderSent bool
	iswritten    bool
}

func (w *doneWriter) WriteHeader(status int) {
	w.isheaderSent = true
	w.ResponseWriter.WriteHeader(status)
}

func (w *doneWriter) Write(b []byte) (int, error) {
	w.iswritten = true
	return w.ResponseWriter.Write(b)
}

func Panic(n http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		dw := &doneWriter{ResponseWriter: rw}
		defer func() {
			if p := recover(); p != nil {

				if !dw.isheaderSent {
					rw.Header().Add("Content-Type", "text/html")
					rw.WriteHeader(http.StatusInternalServerError)
				}

				errstack := debug.Stack()
				fmt.Fprintf(rw, "<h1>Unexpected error occurs</h1>")
				fmt.Fprintf(rw, "<code>%+v<br />", p)
				rw.Write(errstack)
				fmt.Fprint(rw, "</code>")

				return
			}
		}()
		n.ServeHTTP(dw, r)
	})
}
