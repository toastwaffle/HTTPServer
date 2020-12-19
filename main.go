// Binary main runs the server.
package main

import (
	"flag"
	"log"

	"response"
	"router"
	"server"
)

var (
	host = flag.String("host", "", "host for server to listen on")
	port = flag.Int("port", 80, "port for server to listen on")
)

func main() {
	flag.Parse()

	r := router.New(
		router.Constant("dump",
			router.Leaf(response.DumpRequest),
			router.Variable("some_var",
				router.Leaf(response.DumpRequest),
				router.PathLeaf(response.DumpRequest, "path"))),
	)

	s, err := server.New(r, server.Host(*host), server.Port(*port))
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
