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
	fmt.Fprintf(w, "%s%s %s\n", prefix, req.Method, req.URL.String())

	// Headers and body only if verbose mode is enabled
	if verbose {
		for key, values := range req.Header {
			for _, value := range values {
				fmt.Fprintf(w, "%s: %s\n", key, value)
			}
		}
		// Empty line separating headers from body
		fmt.Fprintln(w)

		// Body (if any) - only print when verbose is enabled
		if len(body) > 0 {
			bodyStr := string(body)
			coda := "\n\n"
			if strings.HasSuffix(bodyStr, "\n\n") {
				coda = ""
			} else if strings.HasSuffix(bodyStr, "\n") {
				coda = "\n"
			}
			fmt.Fprint(w, bodyStr+coda)
		}
	}

	return nil
}

// PrintResponse writes an http.Response to stdout or a file
func PrintResponse(resp *http.Response, includeHeaders bool, outputFile string) error {
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	// If outputFile is specified, write to file instead of stdout
	if outputFile != "" {
		// Create the file
		file, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write headers to file if requested
		if includeHeaders {
			_, err = fmt.Fprintln(file, resp.Status)
			if err != nil {
				return err
			}

			err = resp.Header.Write(file)
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(file)
			if err != nil {
				return err
			}
		}

		// Write body to file
		_, err = file.Write(body)
		return err
	}

	// Otherwise write to stdout as before
	if includeHeaders {
		fmt.Println(resp.Status)

		resp.Header.Write(os.Stdout)
		fmt.Println()
	}

	fmt.Println(string(body))

	return nil
}
