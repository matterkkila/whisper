package whisper

import (
	"log"
	"net"
)

type Server struct {
	conn    *net.UDPConn
	parser  *Parser
	workers chan *metricBuffer
}

type metricBuffer struct {
	buf *[]byte
	length int
}

func (s *Server) spawnWorkers(poolSize int) {
	for i := 0; i < poolSize; i++ {
		go func() {
			for mb := range s.workers {
				metric, err := (*s.parser).Parse((*mb.buf)[:mb.length])
				if err == nil {
					s.handleMetric(metric)
				}
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
	s.workers <- &metricBuffer{&buf, n}
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

	workers := make(chan *metricBuffer, 64)
	s := Server{conn, parser, workers}
	s.spawnWorkers(poolSize)

	return &s, nil
}
