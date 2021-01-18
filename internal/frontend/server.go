package frontend

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/zapkub/pakkretqc/internal/fsutil"
)

type Server struct {
	apptemplate map[string]*template.Template
}

func (s *Server) Install(handle func(string, http.Handler)) {

	handle("/login", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s.apptemplate["login"].Execute(rw, nil)
	}))
	handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s.apptemplate["index"].Execute(rw, nil)
	}))
}

func parseTemplates(filename string) *template.Template {
	tmpl, err := template.ParseFiles(
		fsutil.PathFromWebDir("common/base.html"),
		filename,
	)
	if err != nil {
		panic(fmt.Sprintf("BUG: cannot parse template %+v", err))
	}
	return tmpl
}

func New() *Server {

	return &Server{
		apptemplate: map[string]*template.Template{
			"index": parseTemplates(fsutil.PathFromWebDir("index.html")),
			"login": parseTemplates(fsutil.PathFromWebDir("login.html")),
		},
	}
}
