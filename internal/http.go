// Package common provides some utility functions for dealing with HTTP requests and responses
package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"strings"
)

// MakeRequest performs an HTTP request based on the information provided
func MakeRequest(method string, url string, input io.Reader, headersIncluded bool) (*http.Response, error) {
	var req *http.Request
	var headers map[string][]string
	var err error

	method = strings.ToUpper(method)

	if input == nil {
		req, err = http.NewRequest(method, url, nil)
	} else if !headersIncluded {
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

		req.Header = headers
	}

	//return http.DefaultClient.Do(req)

	return http.DefaultTransport.RoundTrip(req)
}

// WriteRequest writes an http.Request to the specified writer
func WriteRequest(w io.Writer, req *http.Request, includeMethod bool, includeUrl bool, includeHeaders bool) error {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		return err
	}

	if includeMethod {
		fmt.Fprintln(w, req.Method)
	}

	if includeUrl {
		fmt.Fprintln(w, req.URL)
	}

	if includeHeaders {
		req.Header.Write(w)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, string(body))

	return nil
}

// WriteResponse writes an http.Response to the specified writer
func WriteResponse(w io.Writer, resp *http.Response, includeHeaders bool, includeStatus bool) error {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return err
	}

	if includeStatus {
		fmt.Fprintln(w, resp.Status)
	}

	if includeHeaders {
		resp.Header.Write(w)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, string(body))

	return nil
}
