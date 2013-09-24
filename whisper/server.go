package whisper

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	conn   *net.UDPConn
	parser *Parser
}

func (s *Server) handleMessage() {
	buf := make([]byte, 1024)

	n, _, err := s.conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	metric, err := (*s.parser).Parse(buf[:n])
	if err != nil {
		log.Printf("bad metric, ignoring. %v", buf[:n])
		return
	}
	s.handleMetric(metric)
}

func (s *Server) handleMetric(metric *Metric) {
	log.Printf("%v", metric)
}

func checkError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Fatal error:%s", err.Error()))
	}
}

func Serve(address string, parser *Parser) {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	s := Server{conn, parser}

	for {
		s.handleMessage()
	}
}
