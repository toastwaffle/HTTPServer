// Package server sets up the server.
package server

import (
	"fmt"
	"log"
	"net"

	"request"
)

type Server struct {
	host string
	port int
}

type ServerOpt func (*Server) error

func Host(host string) ServerOpt {
	return func (s *Server) error {
		s.host = host
		return nil
	}
}

func Port(port int) ServerOpt {
	return func (s *Server) error {
		s.port = port
		return nil
	}
}

func New(opts ...ServerOpt) (*Server, error) {
	s := &Server{
		host: "",
		port: 80,
	}
	for _, o := range opts {
		if err := o(s); err != nil {
			return nil, fmt.Errorf("failed to configure server: %w", err)
		}
	}
	return s, nil
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.host, s.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", s.port, err)
	}
	defer ln.Close()
	log.Printf("Server listening on %s", ln.Addr())
	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("failed to accept connection: %w", err)
		}
		go s.handleConnection(conn)
	}
}

func badRequest(c net.Conn, msg string) {
	fmt.Fprint(c, "HTTP/1.1 400 Bad Request\r\n")
	if msg == "" {
		return
	}
	fmt.Fprintf(c, "Content-Type: text/plain\r\n\r\n%s", msg)
}

func (s *Server) handleConnection(c net.Conn) {
	defer c.Close()

	req, err := request.Parse(c)
	if err != nil {
		badRequest(c, err.Error())
		return
	}

	req.Dump(c)
}
