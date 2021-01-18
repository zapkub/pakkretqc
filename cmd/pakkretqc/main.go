package main

import (
	"log"
	"net/http"

	"github.com/zapkub/pakkretqc/internal/frontend"
	"github.com/zapkub/pakkretqc/internal/pserver"
)

func main() {

	var pserv = pserver.New()
	var frontserv = frontend.New()
	frontserv.Install(pserv.Handle)

	if err := http.ListenAndServe("127.0.0.1:8888", pserv); err != nil {
		log.Fatal(err)
	}

}
