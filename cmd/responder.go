package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/textproto"
	"os"

	"github.com/stilvoid/please/common"
)

type responder struct {
	status          int
	includeHeaders  bool
	includeMethod   bool
	includeUrl      bool
	headersIncluded bool
	listener        net.Listener
	data            io.ReadSeeker
}

func (h responder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer h.listener.Close()

	err := common.WriteRequest(os.Stdout, req, h.includeMethod, h.includeUrl, h.includeHeaders)
	os.Stdout.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	inputReader := bufio.NewReader(h.data)

	if h.headersIncluded {
		if h.data == nil {
			fmt.Println("Error reading headers")
			os.Exit(1)
		}

		// Parse headers from input
		reader := textproto.NewReader(inputReader)
		headers, err := reader.ReadMIMEHeader()

		if err != nil {
			fmt.Println("Error parsing headers:", err)
			os.Exit(1)
		}

		for name, values := range headers {
			for i, value := range values {
				if i == 0 {
					w.Header().Set(name, value)
				} else {
					w.Header().Add(name, value)
				}
			}
		}
	}

	w.WriteHeader(h.status)

	if h.data != nil {
		io.Copy(w, inputReader)
	}
}
