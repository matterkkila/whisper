package main

import (
	"fmt"
	"log"
	"net"
)

const (
	port string = ":1200"
)

func sendMessage(conn net.Conn) {
	_, err := conn.Write([]byte("my.metric.name 1|c 1379700446"))
	if err != nil {
		log.Printf("%v", err)
	}
}

func checkError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Fatal error:%s", err.Error()))
	}
}

func main() {
	conn, err := net.Dial("udp", port)
	checkError(err)

	defer conn.Close()

	for {
		sendMessage(conn)
	}
}
