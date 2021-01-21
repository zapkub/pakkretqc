package frontend

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zapkub/pakkretqc/internal/middleware"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

type defectPage struct {
	Defect *almsdk.Defect `json:"defect"`
}

func (s *Server) defectPageHandler(w http.ResponseWriter, r *http.Request) {
	var (
		page      defectPage
		vars      = mux.Vars(r)
		domain    = vars["domain"]
		project   = vars["project"]
		id        = vars["id"]
		ctx       = r.Context()
		almclient = middleware.MustGetALMClient(ctx)
	)

	deflect, err := almclient.Defect(ctx, domain, project, id)
	if err != nil {
		panic(err)
	}
	page.Defect = deflect

	s.servePage(w, "defect", page)
}
