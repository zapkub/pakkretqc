package frontend

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/zapkub/pakkretqc/internal/conf"
	"github.com/zapkub/pakkretqc/internal/session"
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
		token, _  = r.Cookie(session.CookieKey)
		almclient = almsdk.New(&almsdk.ClientOptions{Endpoint: conf.ALMEndpoint()})
	)
	defer func() {
		s.servePage(w, "login", loginPage)
	}()
	if token == nil {
		if username := r.FormValue("username"); len(username) > 0 && method == "POST" {
			log.Printf("login with %s", username)
			token := base64.URLEncoding.EncodeToString([]byte(username + ":" + r.FormValue("password")))
			err := almclient.Authenticate(ctx, token)
			if err != nil {
				log.Println(err)
			}
			log.Printf("login with %s success", username)
			domains, err := almclient.Domains(ctx)
			if err != nil {
				log.Println(err)
				return
			}
			if domain := r.FormValue("domain"); domain == "" {
				loginPage.Username = username
				loginPage.Domains = domains
			}
			fmt.Printf("%+v", domains)
			var cookietoken http.Cookie
			cookietoken.Path = "/"
			cookietoken.HttpOnly = true
			cookietoken.Name = session.CookieKey
			cookietoken.Value = token
			http.SetCookie(w, &cookietoken)
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
