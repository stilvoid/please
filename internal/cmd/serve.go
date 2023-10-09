package cmd

import (
	"fmt"
	"io"
    "log"
	"net"
	"net/http"
	"os"

	"github.com/pborman/getopt"
)

type server struct {
	status          int
	includeHeaders  bool
	includeMethod   bool
	includeUrl      bool
	headersIncluded bool
	listener        net.Listener
	data            io.ReadSeeker
}

func init() {
	Commands["serve"] = serveCommand
}

func serveHelp() {
	fmt.Println("Usage: please serve <PATH> [<ADDRESS>[:<PORT>]]")
	fmt.Println()
	fmt.Println("Serves the contents of PATH on the specified address and port.")
	fmt.Println("Requested paths will be printed to stdout.")
}

type loggingFileSystem struct {
    http.FileSystem
}

func (fsys loggingFileSystem) Open(path string) (http.File, error) {
    fmt.Println(path)

    return fsys.FileSystem.Open(path)
}

func serveCommand(args []string) {
	opts := getopt.CommandLine

	opts.SetUsage(serveHelp)

	// Deal with flags and get the path (and URL)
	opts.Parse(args)
	if opts.NArgs() < 1 {
		getopt.Usage()
		os.Exit(1)
	}

	path := opts.Arg(0)

	address := "0.0.0.0:8000"

	if opts.NArgs() >= 2 {
		address = opts.Arg(1)
	}

    fmt.Println("Listening on", address)

    fsys := loggingFileSystem{http.Dir(path)}
    log.Fatal(http.ListenAndServe(address, http.FileServer(fsys)))
}
