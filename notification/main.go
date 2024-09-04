package main

import (
	"log"

	"github.com/dilshodforever/5-oyimtixon/server"
)

func main() {
	lis:= server.Connection()
	log.Printf("%v", lis.Addr())
}
