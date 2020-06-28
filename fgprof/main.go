package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/felixge/fgprof"
)

func main() {
	http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
	log.Println(http.ListenAndServe(":6060", nil))
}
