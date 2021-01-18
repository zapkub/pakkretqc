package main

import (
	"log"
	"net/http"

	"github.com/zapkub/pakkretqc/internal/conf"
	"github.com/zapkub/pakkretqc/internal/frontend"
	"github.com/zapkub/pakkretqc/internal/pserver"
	"github.com/zapkub/pakkretqc/pkg/almsdk"
)

func main() {

	var almclient = almsdk.New(&almsdk.ClientOptions{
		Endpoint: conf.ALMEndpoint(),
	})
	var pserv = pserver.New()
	var frontserv = frontend.New(almclient)
	frontserv.Install(pserv.Handle)

	if err := http.ListenAndServe("127.0.0.1:8888", pserv); err != nil {
		log.Fatal(err)
	}

}
