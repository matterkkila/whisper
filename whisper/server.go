package whisper

import (
	"log"
	"net"
)

type Server struct {
	conn    *net.UDPConn
	parser  *Parser
	workers chan *Metric
}

func (s *Server) spawnWorkers(poolSize int) {
	for i := 0; i < poolSize; i++ {
		go func() {
			for metric := range s.workers {
				s.handleMetric(metric)
			}
		}()
	}
}

func (s *Server) handleMessage() {
	buf := make([]byte, 4096)
	n, _, err := s.conn.ReadFromUDP(buf[0:])
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	metric, err := (*s.parser).Parse(buf[:n])
	if err != nil {
		log.Printf("bad metric, ignoring. %v", buf[:n])
		return
	}
	s.workers <- metric
}

func (s *Server) handleMetric(metric *Metric) {
	log.Printf("%v", metric)
}

func (s *Server) Serve() {
	for {
		s.handleMessage()
	}
	close(s.workers)
}

func NewServer(address string, parser *Parser, poolSize int) (*Server, error) {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, err
	}

	workers := make(chan *Metric, 64)
	s := Server{conn, parser, workers}
	s.spawnWorkers(poolSize)

	return &s, nil
}
