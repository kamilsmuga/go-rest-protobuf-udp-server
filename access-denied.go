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

type Stats struct {
	app_sha256 string
	ip         int64
	count      uint64
	goodIps    []int64
	badIps     []int64
}

func EventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//TODO Return statistics per app_sha256
		/* Format:
					{
		  				'count':NUMBER_OF_EVENTS,
		  				'good_ips':LIST_OF_GOOD_IPS,
		  				'bad_ips':LIST_OF_BAD_IPS
					}
		*/
		log.Println("GET: " + r.RequestURI)
	} else if r.Method == "DELETE" {
		//TODO Purge statistics for app_sha256
		log.Println("DELETE: " + r.RequestURI)
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
	http.HandleFunc("/events/", EventsHandler)
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
	//TODO Decode protobuf message
	//TODO Validate correct IP range
	//TODO Store total count, good and bad ips
	fmt.Println(string(buf[0:rlen]))
}
