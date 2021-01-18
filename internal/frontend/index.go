package frontend

import "net/http"

func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {

	s.servePage(w, "index", nil)
}
