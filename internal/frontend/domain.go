package frontend

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zapkub/pakkretqc/internal/middleware"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

type domainPage struct {
	Projects []almsdk.Projects `json:"projects"`
	Domain   string            `json:"domain"`
}

func (s *Server) domainHandler(w http.ResponseWriter, r *http.Request) {
	var (
		page      domainPage
		ctx       = r.Context()
		almclient = middleware.MustGetALMClient(ctx)
	)

	vars := mux.Vars(r)
	projects, err := almclient.Projects(ctx, vars["domain"])
	if err != nil {
		log.Printf("ERROR: %+v", err)
		return
	}
	page.Projects = projects
	page.Domain = vars["domain"]
	s.servePage(w, "domain", page)

}
