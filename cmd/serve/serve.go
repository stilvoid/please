package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var path string
var address string
var port int

func init() {
	Cmd.Flags().StringVarP(&path, "dir", "d", "./", "Directory to serve")
	Cmd.Flags().StringVarP(&address, "address", "a", "", "Address to listen on")
	Cmd.Flags().IntVarP(&port, "port", "p", 8000, "Post to listen on")
}

func serveHelp() {
	fmt.Println("Usage: please serve <PATH> [<ADDRESS>[:<PORT>]]")
	fmt.Println()
	fmt.Println("Serves the contents of PATH on the specified address and port.")
	fmt.Println("Requested paths will be printed to stdout.")
}

var Cmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the contents of a directory through a simple web server",
	Run: func(cmd *cobra.Command, args []string) {
		address = fmt.Sprintf("%s:%d", address, port)

		fmt.Println("Listening on", address)

		fsys := loggingFileSystem{http.Dir(path)}

		cobra.CheckErr(http.ListenAndServe(address, http.FileServer(fsys)))
	},
}

type loggingFileSystem struct {
	http.FileSystem
}

func (fsys loggingFileSystem) Open(path string) (http.File, error) {
	log.Println(path)

	return fsys.FileSystem.Open(path)
}
