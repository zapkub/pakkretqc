package frontend

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

type projectPage struct {
	Total    int               `json:"total"`
	Domain   string            `json:"domain"`
	Project  string            `json:"project"`
	Defects  []*almsdk.Deflect `json:"defects"`
	Username string            `json:"username"`
}

func (s *Server) projectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		page    projectPage
		ctx     = almsdk.AppendSessionCookieToContext(r.Context(), r)
		domain  = vars["domain"]
		project = vars["project"]
	)

	defer func() {
		s.servePage(w, "project", page)
	}()
	deflect, total, err := s.almclient.Deflects(ctx, domain, project, 50, 0, "-creation-time")
	if err != nil {
		log.Printf("ERROR: %+v", err)
		return
	}
	page.Project = project
	page.Domain = domain
	page.Defects = deflect
	page.Total = total
	page.Username = UserName(r)

}
