package singleton

import "fmt"

type Server struct {
	params string
}

func NewServer(params string) *Server {
	return &Server{params: params}
}

var server *Server

func GetInstance() *Server {
	if server == nil {
		server = NewServer("params")
		fmt.Println("Creating new Server")
	} else {
		fmt.Println("Using already created server", &server)
	}
	return server
}
