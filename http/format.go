// Package http provides some utility functions for dealing with HTTP requests and responses
package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// FormatRequest produces a string representation of an http request
func FormatRequest(req *http.Request, o Options) (string, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	req.Body.Close()

	out := &strings.Builder{}

	if o.IncludeMethod {
		out.WriteString(req.Method)

		if o.IncludePath {
			out.WriteString(" ")
		} else {
			out.WriteString("\n")
		}
	}

	if o.IncludePath {
		out.WriteString(req.URL.String())
		out.WriteString("\n")
	}

	if o.IncludeHeaders {
		req.Header.Write(out)
		out.WriteString("\n")
	}

	out.WriteString(string(body))

	return out.String(), nil
}

// FormatResponse produces a string representation of an http response
func FormatResponse(resp *http.Response, o Options) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()

	out := &strings.Builder{}

	if o.IncludeStatus {
		out.WriteString(resp.Status)
		out.WriteString("\n")
	}

	if o.IncludeHeaders {
		resp.Header.Write(out)
		out.WriteString("\n")
	}

	out.WriteString(string(body))

	return out.String(), nil
}
