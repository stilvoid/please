package main

import (
	"github.com/stilvoid/please/cmd"
)

func main() {
	/*
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(os.Stderr, "%s\n", fmt.Sprint(r))
				os.Exit(1)
			}

			os.Exit(0)
		}()
	*/

	cmd.Execute()
}
