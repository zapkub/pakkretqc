package frontend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zapkub/pakkretqc/internal/middleware"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

type defectPage struct {
	Defect     *almsdk.Defect       `json:"defect"`
	Attachment []*almsdk.Attachment `json:"attachment"`
	Project    string               `json:"project"`
	Domain     string               `json:"domain"`
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

	attachment, err := almclient.Attachments(ctx, domain, project, fmt.Sprintf("parent-id = %s ; parent-type = '%s'", id, "defect"), 10, 0)
	if err != nil {
		panic(err)
	}
	page.Attachment = attachment
	page.Domain = domain
	page.Project = project

	s.servePage(w, "defect", page)
}

func (s *Server) attachmentDownloadHandler(w http.ResponseWriter, r *http.Request) {

	var (
		ctx       = r.Context()
		vars      = mux.Vars(r)
		domain    = vars["domain"]
		project   = vars["project"]
		id        = vars["id"]
		almclient = middleware.MustGetALMClient(ctx)
	)

	err := almclient.Attachment(ctx, domain, project, id, w)
	if err != nil {
		log.Printf("ERROR: %+v", err)
	}

}
