package main

import (
	"github.com/whisper/whisper"
)

const (
	port string = ":1200"
)

func main() {
	parser := whisper.Parser(whisper.TextParser{})
	whisper.Serve(port, &parser)
}
