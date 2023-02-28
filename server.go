package main

type Server struct {
	listenAddr int
	message    string
}

func NewServer(addr int) *Server {
	return &Server{
		listenAddr: addr,
	}
}
