package main

import (
	"fmt"
	"github.com/whisper/whisper"
)

const (
	port     string = ":1200"
	poolSize int    = 4
)

func checkError(err error) {
	if err != nil {
		panic(fmt.Sprintf("Fatal error:%s", err.Error()))
	}
}

func main() {
	parser := whisper.Parser(whisper.TextParser{})
	server, err := whisper.NewServer(port, &parser, poolSize)
	checkError(err)
	server.Serve()
}
