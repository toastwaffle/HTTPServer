// Binary main runs the server.
package main

import (
	"flag"
	"log"

	"server"
)

var (
	host = flag.String("host", "", "host for server to listen on")
	port = flag.Int("port", 80, "port for server to listen on")
)

func main() {
	flag.Parse()

	s, err := server.New(server.Host(*host), server.Port(*port))
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
