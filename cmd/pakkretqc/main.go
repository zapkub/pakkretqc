package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/zapkub/pakkretqc/internal/frontend"
	"github.com/zapkub/pakkretqc/internal/middleware"
	"github.com/zapkub/pakkretqc/internal/pserver"
)

func main() {

	var pserv = pserver.New()
	var frontserv = frontend.New()
	frontserv.Install(pserv.Handle)

	mw := middleware.Chain(
		middleware.Panic,
		middleware.Session,
		middleware.ALMClient,
	)

	var addr = "0.0.0.0:8888"
	fmt.Printf("Application started.\naddr: http://%s\nlocal: http://%s\n", addr, strings.Replace(addr, "0.0.0.0", "localhost", -1))
	if err := http.ListenAndServe(addr, mw(pserv)); err != nil {
		log.Fatal(err)
	}

}
