package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var (
	host    = flag.String("host", "127.0.0.1", "host IP")
	udpPort = flag.String("udp", "8080", "udp port")
	tcpPort = flag.String("tcp", "8081", "tcp port")
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
	flag.Parse()
	go serveUdp()
	serveTcp()
}

func serveTcp() {
	http.HandleFunc("/events", EventsHandler)
	log.Printf("Serving TCP requests on: %s:%s\n", *host, *tcpPort)
	err := http.ListenAndServe(*host+":"+*tcpPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func serveUdp() {
	log.Printf("Serving UDP requests on: %s:%s\n", *host, *udpPort)
	addr, _ := net.ResolveUDPAddr("udp", ":"+*udpPort)
	sock, _ := net.ListenUDP("udp", addr)
	sock.SetReadBuffer(1048576)
	for {
		buf := make([]byte, 1024)
		rlen, _, err := sock.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		go handlePacket(buf, rlen)
	}
}

func handlePacket(buf []byte, rlen int) {
	fmt.Println(string(buf[0:rlen]))
}
