package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/textproto"
	"os"
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

		var body []byte
		body, err = ioutil.ReadAll(inputReader)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(body))

		req.Header = headers
	}

	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))

	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

// PrintRequest writes an http.Request to stdout
func PrintRequest(req *http.Request, includeHeaders bool) error {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		return err
	}

	fmt.Printf("%s %s\n", req.Method, req.URL)

	if includeHeaders {
		req.Header.Write(os.Stdout)
		fmt.Println()
	}

	fmt.Println(string(body))

	return nil
}

// PrintResponse writes an http.Response to stdout
func PrintResponse(resp *http.Response, includeHeaders bool) error {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	if includeHeaders {
		fmt.Println(resp.Status)

		resp.Header.Write(os.Stdout)
		fmt.Println()
	}

	fmt.Println(string(body))

	return nil
}
