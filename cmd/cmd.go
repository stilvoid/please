/*
Package cmd houses the command line functionality of please.
*/
package cmd

// Command represents a callable command
type Command func([]string)

// Commands stores a mapping of command names to their functions
var Commands = make(map[string]Command)

// Aliases stores a mapping of command aliases to command names
var Aliases = make(map[string]string)
