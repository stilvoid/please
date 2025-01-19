package main

import (
	"log"

	"github.com/stilvoid/please/cmd/please"
)

func main() {
	log.SetFlags(0)

	if err := please.Cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
