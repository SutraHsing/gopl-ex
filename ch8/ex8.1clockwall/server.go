package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// get port varname and its value
	port := flag.String("port", "8010", "the port server")
	flag.Parse()
	if *port != "8010" && *port != "8020" && *port != "8030" {
		log.Fatalf("port %s not supported.\n", *port)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Print(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn, *port)
	}

}

func handleConn(conn net.Conn, port string) {
	defer conn.Close()
	// different time zone
	var zone string
	switch port {
	case "8010":
		zone = "US/Eastern"
	case "8020":
		zone = "Asia/Tokyo"
	case "8030":
		zone = "Europe/London"
	default:
		log.Printf("unsupported zone: %s\n", zone)
	}
	location, err := time.LoadLocation(zone)
	if err != nil {
		log.Print(err)
	}
	t := time.Now().In(location).Format("15:04:05")
	_, err = io.WriteString(conn, t)
	if err != nil {
		log.Print(err)
	}
}
