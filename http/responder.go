package http

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
)

type Responder struct {
	Data    io.ReadSeeker
	Stop    chan bool
	Options Options
	Status  int
}

func (r Responder) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqString, err := FormatRequest(req, r.Options)
	if err != nil {
		panic(err)
	}
	fmt.Println(reqString)

	var headers textproto.MIMEHeader
	var input []byte

	if r.Data != nil {
		inputReader := bufio.NewReader(r.Data)

		// Parse headers from input
		if r.Options.HeadersIncluded {
			reader := textproto.NewReader(inputReader)
			headers, err = reader.ReadMIMEHeader()
			if err != nil {
				panic(fmt.Errorf("Error parsing headers: %s", err.Error()))
			}
		}

		// Read input
		input, err = ioutil.ReadAll(inputReader)
		if err != nil {
			panic(fmt.Errorf("Unable to read request body: %s", err.Error()))
		}
	}

	// Write headers
	for name, values := range headers {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// Write status
	w.WriteHeader(r.Status)

	// Write body
	w.Write(input)

	// Close the connection
	// TODO: Allow this to optionally go until closed
	r.Stop <- true
}
