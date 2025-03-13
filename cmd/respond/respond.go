package respond

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/textproto"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/stilvoid/please/internal"
)

var headersIncluded bool
var verbose bool
var inputFile string
var outputFile string
var address string
var port int
var status int
var keepAlive bool

func init() {
	Cmd.Flags().BoolVarP(&headersIncluded, "include-headers", "i", false, "Read headers from the response body")
	Cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show full request details including headers and body")
	Cmd.Flags().StringVarP(&inputFile, "data", "d", "", "Filename to read the response body from. Omit for stdin.")
	Cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Filename to write the request to. Omit for stdout.")
	Cmd.Flags().StringVarP(&address, "address", "a", "", "Address to listen on")
	Cmd.Flags().IntVarP(&port, "port", "p", 8000, "Port to listen on")
	Cmd.Flags().IntVarP(&status, "status", "s", 200, "Status code to respond with")
	Cmd.Flags().BoolVarP(&keepAlive, "keep-alive", "k", false, "Keep server running to handle multiple requests")
}

var Cmd = &cobra.Command{
	Use:   "respond",
	Short: "Listen for HTTP requests and respond to them",
	Long:  "Listen for HTTP requests and respond to them. By default, responds to a single request and exits. Use --keep-alive to keep the server running.",
	Run: func(cmd *cobra.Command, args []string) {
		address = fmt.Sprintf("%s:%d", address, port)

		// Read the response data once at startup
		var responseData []byte
		var err error

		if inputFile == "" {
			// If no input file is specified, try to read from stdin
			stdinData, err := internal.StdinOrNothing()
			if err != nil {
				cobra.CheckErr(err)
			}
			if stdinData != nil {
				responseData, err = ioutil.ReadAll(stdinData)
				cobra.CheckErr(err)
			}
		} else {
			// Read from the specified file
			responseData, err = os.ReadFile(inputFile)
			cobra.CheckErr(err)
		}

		var outputWriter io.Writer

		if outputFile == "" {
			outputWriter = os.Stdout
		} else {
			f, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			cobra.CheckErr(err)

			defer f.Close()
			outputWriter = f
		}

		handler := responder{
			status:          status,
			verbose:         verbose,
			headersIncluded: headersIncluded,
			responseData:    responseData,
			keepAlive:       keepAlive,
			outputWriter:    outputWriter,
		}

		listener, err := net.Listen("tcp", address)
		cobra.CheckErr(err)

		server := &http.Server{Addr: address, Handler: &handler}

		// Channel for signaling server shutdown
		ch = make(chan bool, 1)

		// Set up signal handling for graceful shutdown
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			err := server.Serve(listener)
			if err != nil && err.Error() != "http: Server closed" {
				cobra.CheckErr(err)
			}
		}()

		fmt.Println("Listening on", address)
		if keepAlive {
			fmt.Println("Server will keep running. Press Ctrl+C to stop.")
		} else {
			fmt.Println("Server will exit after the first request.")
		}

		// Wait for either a shutdown signal from the handler or an OS signal
		select {
		case <-ch:
			// Normal shutdown after request handling
		case <-sigChan:
			// Shutdown due to OS signal
			fmt.Println("\nShutting down server...")
		}

		server.Shutdown(context.Background())
	},
}

var ch chan bool

type responder struct {
	status          int
	verbose         bool
	headersIncluded bool
	responseData    []byte
	keepAlive       bool
	requestCount    int
	outputWriter    io.Writer
}

func (h *responder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.requestCount++
	requestNum := h.requestCount

	prefix := ""
	if h.keepAlive {
		prefix = fmt.Sprintf("Request #%d: ", requestNum)
	}

	err := internal.WriteRequest(h.outputWriter, req, h.verbose, prefix)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing request: %v\n", err)
	}

	// Set up the response
	if h.headersIncluded && len(h.responseData) > 0 {
		// Parse headers from the response data
		reader := textproto.NewReader(bufio.NewReader(bytes.NewReader(h.responseData)))
		headers, err := reader.ReadMIMEHeader()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing headers: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Set the headers in the response
		for name, values := range headers {
			for i, value := range values {
				if i == 0 {
					w.Header().Set(name, value)
				} else {
					w.Header().Add(name, value)
				}
			}
		}

		// Find where the headers end and the body begins
		bodyStart := 0
		for i := 0; i < len(h.responseData); i++ {
			if i+3 < len(h.responseData) &&
				h.responseData[i] == '\r' && h.responseData[i+1] == '\n' &&
				h.responseData[i+2] == '\r' && h.responseData[i+3] == '\n' {
				bodyStart = i + 4
				break
			}
			if i+1 < len(h.responseData) &&
				h.responseData[i] == '\n' && h.responseData[i+1] == '\n' {
				bodyStart = i + 2
				break
			}
		}

		// Write the status code and body
		w.WriteHeader(h.status)
		w.Write(h.responseData[bodyStart:])
	} else {
		// No headers to parse, just write the status and body
		w.WriteHeader(h.status)
		if h.responseData != nil {
			w.Write(h.responseData)
		}
	}

	// If not in keep-alive mode, signal to shut down the server after this request
	if !h.keepAlive {
		go func() {
			ch <- true
		}()
	}
}
