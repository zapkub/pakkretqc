package pserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zapkub/pakkretqc/internal/fsutil"
)

type Transport struct {
	mux *mux.Router
}

func New() *Transport {
	r := mux.NewRouter()
	r.PathPrefix("/dist").Handler(
		http.StripPrefix("/dist",
			http.FileServer(http.Dir(fsutil.PathFromWebDir("dist"))),
		),
	)
	return &Transport{
		mux: r,
	}
}

func (h *Transport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *Transport) Handle(path string, handler http.Handler) {
	h.mux.Path(path).Handler(handler)
}
