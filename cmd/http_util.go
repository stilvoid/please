package cmd

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

func createRequest(method string, url string, input io.Reader, headers_included bool) *http.Request {
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
			fmt.Println("Error parsing headers:", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(input_reader)

		if err != nil {
			fmt.Println("Error reading input:", err)
			os.Exit(1)
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	if input != nil && headers_included {
		for name, values := range headers {
			for i, value := range values {
				if i == 0 {
					req.Header.Set(name, value)
				} else {
					req.Header.Add(name, value)
				}
			}
		}
	}

	return req
}

func getResponse(req *http.Request) *http.Response {
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error contacting host:", err)
		os.Exit(1)
	}

	return resp
}

func printRequest(req *http.Request, include_method bool, include_url bool, include_headers bool) {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if include_method {
		fmt.Println(req.Method)
	}

	if include_url {
		fmt.Println(req.URL)
	}

	if include_headers {
		req.Header.Write(os.Stdout)
		fmt.Println()
	}

	fmt.Println(string(body))
}

func printResponse(resp *http.Response, include_headers bool, include_status bool) {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if include_status {
		fmt.Println(resp.StatusCode)
	}

	if include_headers {
		resp.Header.Write(os.Stdout)
		fmt.Println()
	}

	fmt.Println(string(body))
}
