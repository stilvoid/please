package util

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

func CreateRequest(method string, url string, input io.Reader, headers_included bool) (*http.Request, error) {
	var req *http.Request
	var headers map[string][]string
	var err error

	method = strings.ToUpper(method)

	if input == nil {
		req, err = http.NewRequest(method, url, nil)
	} else if !headers_included {
		req, err = http.NewRequest(method, url, input)
	} else {
		input_reader := bufio.NewReader(input)

		reader := textproto.NewReader(input_reader)

		headers, err = reader.ReadMIMEHeader()

		if err != nil {
			return nil, err
		}

		body, err := ioutil.ReadAll(input_reader)

		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(body))

		req.Header = headers
	}

	return req, err
}

func GetResponse(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}

func PrintRequest(w io.Writer, req *http.Request, include_method bool, include_url bool, include_headers bool) error {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		return err
	}

	if include_method {
		fmt.Fprintln(w, req.Method)
	}

	if include_url {
		fmt.Fprintln(w, req.URL)
	}

	if include_headers {
		req.Header.Write(w)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, string(body))

	return nil
}

func PrintResponse(w io.Writer, resp *http.Response, include_headers bool, include_status bool) error {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return err
	}

	if include_status {
		fmt.Fprintln(w, resp.StatusCode)
	}

	if include_headers {
		resp.Header.Write(w)
		fmt.Fprintln(w)
	}

	fmt.Fprintln(w, string(body))

	return nil
}
