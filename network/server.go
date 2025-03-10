package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(serverOpts ServerOpts) *Server {
	return &Server{
		ServerOpts: serverOpts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}),
	}
}

func (s *Server) Start() {
	s.initTransports()

	ticker := time.NewTicker(5 * time.Second)
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Println("do stuff every x second")

		}

	}
	fmt.Println("Server shutdown")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
