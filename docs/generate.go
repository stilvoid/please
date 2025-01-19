package main

//go:generate go run generate.go

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/stilvoid/please/cmd/please"
)

func main() {
	if err := doc.GenMarkdownTree(please.Cmd, "./"); err != nil {
		log.Fatal(err)
	}

}
