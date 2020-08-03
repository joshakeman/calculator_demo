package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":80", "http service address")

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.Handle("/", http.FileServer(http.Dir("./build")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
