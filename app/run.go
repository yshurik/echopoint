package app

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var addr = flag.String("addr", ":8080", "echo point listen address")
var path = flag.String("path", "/echo", "echo point address path")

func Run() error {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc(*path, func(w http.ResponseWriter, r *http.Request) {
		serve(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return err
	}
	return nil
}
