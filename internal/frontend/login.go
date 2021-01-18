package frontend

import (
	"log"
	"net/http"
	"path"

	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

type loginPage struct {
	Username string            `json:"username"`
	Domains  []*almsdk.Domains `json:"domains"`
}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	method := r.Method
	var (
		loginPage loginPage
		token, _  = r.Cookie("LWSSO_COOKIE_KEY")
	)
	defer func() {
		s.servePage(w, "login", loginPage)
	}()
	if token == nil {
		if username := r.FormValue("username"); len(username) > 0 && method == "POST" {
			log.Printf("login with %s", username)
			sessioncookies, err := s.almclient.Authenticate(ctx, username, r.FormValue("password"))
			if err != nil {
				log.Println(err)
			}
			for _, cookie := range sessioncookies {
				cookie.Secure = false
				http.SetCookie(w, cookie)
				r.AddCookie(cookie)
			}
			ctx = almsdk.AppendSessionCookieToContext(ctx, r)
			domains, err := s.almclient.Domains(ctx)
			if domain := r.FormValue("domain"); domain == "" {
				loginPage.Username = username
				loginPage.Domains = domains
			}
			var cookieusername http.Cookie
			cookieusername.Path = "/"
			cookieusername.Name = "username"
			cookieusername.Value = username
			http.SetCookie(w, &cookieusername)
		}
		return
	}

	if token.Value != "" && method == "POST" {
		switch r.FormValue("action") {
		case "cancel":
			for _, cook := range r.Cookies() {
				cook.MaxAge = -1
				http.SetCookie(w, cook)
			}
			s.servePage(w, "login", loginPage)
			return
		case "proceed":
			currentDomain := r.FormValue("currentDomain")
			http.Redirect(w, r, path.Join("/", "domains", currentDomain), http.StatusTemporaryRedirect)
			return
		}
	}

	token.MaxAge = -1
	http.SetCookie(w, token)

}
