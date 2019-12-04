// Package http provides some utility functions for dealing with HTTP requests and responses
package http

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/stilvoid/please/version"
)

// MakeRequest performs an HTTP request based on the information provided
func MakeRequest(method string, url string, input io.Reader, o Options) (*http.Response, error) {
	var req *http.Request
	var headers = make(map[string][]string)
	var err error

	method = strings.ToUpper(method)

	if input == nil || !o.HeadersIncluded {
		req, err = http.NewRequest(method, url, input)
	} else {
		inputReader := bufio.NewReader(input)

		reader := textproto.NewReader(inputReader)

		headers, err = reader.ReadMIMEHeader()
		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(inputReader)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	}

	// Set default UA
	if _, ok := headers["User-Agent"]; !ok {
		headers["User-Agent"] = []string{version.String()}
	}

	req.Header = headers

	return http.DefaultClient.Do(req)

	//return http.DefaultTransport.RoundTrip(req)
}
