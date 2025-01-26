package serve

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var address string
var port int

func init() {
	Cmd.Flags().StringVarP(&address, "address", "a", "", "Address to listen on")
	Cmd.Flags().IntVarP(&port, "port", "p", 8000, "Post to listen on")
}

var Cmd = &cobra.Command{
	Use:   "serve (PATH)",
	Short: "Serve the contents of PATH (current directory if omitted) through a simple web server",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "./"
		if len(args) == 1 {
			path = args[0]
		}

		address = fmt.Sprintf("%s:%d", address, port)

		fmt.Printf("Serving %s on %s\n", path, address)

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
