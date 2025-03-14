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
func MakeRequest(method string, url string, input io.Reader, headersIncluded bool, customHeaders map[string][]string) (*http.Response, error) {
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

	if err != nil {
		return nil, err
	}

	// Add custom headers from command line
	for name, values := range customHeaders {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	// Add User-Agent header last to ensure it's not overridden
	req.Header.Add("User-Agent", fmt.Sprintf("%s/%s", Name, Version))

	return http.DefaultClient.Do(req)
}

// WriteRequest writes an http.Request to a writer
func WriteRequest(w io.Writer, req *http.Request, verbose bool, prefix string) error {
	body, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		return err
	}

	// Create a new ReadCloser with the same body for future use
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	// First line: HTTP method and path
	_, err = fmt.Fprintf(w, "%s%s %s\n", prefix, req.Method, req.URL.String())
	if err != nil {
		return err
	}

	// Headers and body only if verbose mode is enabled
	if verbose {
		for key, values := range req.Header {
			for _, value := range values {
				_, err = fmt.Fprintf(w, "%s: %s\n", key, value)
				if err != nil {
					return err
				}
			}
		}
		// Empty line separating headers from body
		_, err = fmt.Fprintln(w)
		if err != nil {
			return err
		}

		// Body (if any) - only print when verbose is enabled
		if len(body) > 0 {
			bodyStr := string(body)
			coda := "\n\n"
			if strings.HasSuffix(bodyStr, "\n\n") {
				coda = ""
			} else if strings.HasSuffix(bodyStr, "\n") {
				coda = "\n"
			}
			_, err = fmt.Fprint(w, bodyStr+coda)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// WriteResponse writes an http.Response to a writer
func WriteResponse(w io.Writer, resp *http.Response, verbose bool) error {
	var err error

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	// Otherwise write to stdout as before
	if verbose {
		_, err = fmt.Fprintln(w, resp.Status)
		if err != nil {
			return err
		}

		err = resp.Header.Write(w)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(w)
		if err != nil {
			return err
		}
	}

	_, err = w.Write(body)

	return err
}
