package frontend

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/zapkub/pakkretqc/internal/fsutil"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

type Server struct {
	almclient   *almsdk.Client
	apptemplate map[string]*template.Template
}

func (s *Server) Install(handle func(string, http.Handler)) {
	handle("/login", http.HandlerFunc(s.loginHandler))
	handle("/domains/{domain}/projects/{project}", http.HandlerFunc(s.projectHandler))
	handle("/domains/{domain}", http.HandlerFunc(s.domainHandler))
	handle("/", http.HandlerFunc(s.indexHandler))
}

func parseTemplates(filename string) *template.Template {
	var err error
	tmpl, err := template.
		New("base.html").
		Funcs(template.FuncMap{
			"toJSON": EncodeJSON,
		}).
		ParseFiles(
			fsutil.PathFromWebDir("common/base.html"),
			filename,
		)
	if err != nil {
		panic(fmt.Sprintf("BUG: cannot parse template %+v", err))
	}
	return tmpl
}

func (s *Server) servePage(w http.ResponseWriter, pagename string, page interface{}) {
	log.Println("serve page", pagename)

	err := s.apptemplate[pagename].Execute(w, page)
	if err != nil {
		panic(fmt.Sprintf("BUG: cannot serve page %+v", err))
	}
}

func New(almclient *almsdk.Client) *Server {

	return &Server{
		almclient: almclient,
		apptemplate: map[string]*template.Template{
			"index":   parseTemplates(fsutil.PathFromWebDir("index.html")),
			"login":   parseTemplates(fsutil.PathFromWebDir("login.html")),
			"domain":  parseTemplates(fsutil.PathFromWebDir("domain.html")),
			"project": parseTemplates(fsutil.PathFromWebDir("project.html")),
		},
	}
}

func UserName(r *http.Request) string {
	if usernamecookie, err := r.Cookie("username"); err == nil {
		return usernamecookie.Value
	}
	return ""
}
