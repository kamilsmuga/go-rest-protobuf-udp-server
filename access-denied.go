package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	address = flag.String("address", "127.0.0.1:8080", "bind host:port")
)

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("GET")
	} else if r.Method == "DELETE" {
		log.Println("DELETE")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/events", EventsHandler)
	log.Printf("Serving requests on: %s \n", *address)
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
