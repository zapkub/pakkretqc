package frontend

import (
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zapkub/pakkretqc/internal/middleware"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

type projectPage struct {
	Total    int              `json:"total"`
	Domain   string           `json:"domain"`
	Project  string           `json:"project"`
	Defects  []*defectWithURL `json:"defects"`
	Username string           `json:"username"`
}

type defectWithURL struct {
	*almsdk.Defect
	URL string `json:"url"`
}

func (s *Server) projectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		page      projectPage
		ctx       = r.Context()
		domain    = vars["domain"]
		project   = vars["project"]
		almclient = middleware.MustGetALMClient(ctx)
	)

	page.Project = project
	page.Domain = domain
	defer func() {
		s.servePage(w, "project", page)
	}()
	deflect, total, err := almclient.Defects(ctx, domain, project, 50, 0, "-creation-time", r.FormValue("query"))
	if err != nil {
		log.Printf("ERROR: %+v", err)
		return
	}
	for _, d := range deflect {
		page.Defects = append(page.Defects, &defectWithURL{Defect: d, URL: path.Join("/domains", domain, "/projects", project, "/defects", strconv.Itoa(d.ID))})
	}
	page.Total = total
	page.Username = UserName(r)

}
