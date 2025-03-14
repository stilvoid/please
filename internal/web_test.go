package internal_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stilvoid/please/internal"
)

func TestMakeRequest(t *testing.T) {
	// Create a test server that doesn't block
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("test response"))
	}))
	defer ts.Close()

	// Test GET request
	resp, err := internal.MakeRequest("GET", ts.URL, nil, false, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Test POST request with body
	body := strings.NewReader("test body")
	resp, err = internal.MakeRequest("POST", ts.URL, body, false, map[string][]string{
		"Content-Type": {"text/plain"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Test with headers included in the body
	headerAndBody := strings.NewReader("Content-Type: application/json\r\n\r\n{\"test\":\"data\"}")
	resp, err = internal.MakeRequest("POST", ts.URL, headerAndBody, true, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Test with invalid URL
	_, err = internal.MakeRequest("GET", "http://invalid-url-that-doesnt-exist.example", nil, false, nil)
	if err == nil {
		t.Error("Expected error for invalid URL")
	}
}

func TestWriteRequest(t *testing.T) {
	// Test with empty body
	req, _ := http.NewRequest("GET", "http://example.com", io.NopCloser(strings.NewReader("")))
	req.Header.Set("Content-Type", "text/plain")

	var buf bytes.Buffer
	err := internal.WriteRequest(&buf, req, false, "> ")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "> GET http://example.com") {
		t.Error("Expected request line in output")
	}
	if strings.Contains(output, "Content-Type: text/plain") {
		t.Error("Headers should not be included when verbose is false")
	}

	// Test with body and verbose mode
	buf.Reset()
	bodyContent := "test body"
	req, _ = http.NewRequest("POST", "http://example.com", io.NopCloser(strings.NewReader(bodyContent)))
	req.Header.Set("Content-Type", "text/plain")

	err = internal.WriteRequest(&buf, req, true, "> ")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	output = buf.String()
	if !strings.Contains(output, "> POST http://example.com") {
		t.Error("Expected request line in output")
	}
	if !strings.Contains(output, "Content-Type: text/plain") {
		t.Error("Expected headers in output when verbose is true")
	}
	if !strings.Contains(output, bodyContent) {
		t.Error("Expected body in output when verbose is true")
	}
}

func TestWriteResponse(t *testing.T) {
	// Create a test response
	resp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"text/plain"},
		},
		Body: io.NopCloser(strings.NewReader("test response")),
	}

	// Test with verbose mode
	var buf bytes.Buffer
	err := internal.WriteResponse(&buf, resp, true)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "200 OK") {
		t.Error("Expected status in output")
	}
	if !strings.Contains(output, "Content-Type: text/plain") {
		t.Error("Expected headers in output")
	}
	if !strings.Contains(output, "test response") {
		t.Error("Expected body in output")
	}

	// Test without verbose mode
	buf.Reset()
	resp = &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Content-Type": []string{"text/plain"},
		},
		Body: io.NopCloser(strings.NewReader("test response")),
	}

	err = internal.WriteResponse(&buf, resp, false)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	output = buf.String()
	if strings.Contains(output, "200 OK") {
		t.Error("Status should not be included when verbose is false")
	}
	if strings.Contains(output, "Content-Type: text/plain") {
		t.Error("Headers should not be included when verbose is false")
	}
	if !strings.Contains(output, "test response") {
		t.Error("Expected body in output")
	}

	// Test error handling when reading body fails
	badResp := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       io.NopCloser(errorReader{}),
	}

	err = internal.WriteResponse(&buf, badResp, true)
	if err == nil {
		t.Error("Expected error when reading body fails")
	}
}

// errorReader is a mock reader that always returns an error
type errorReader struct{}

func (e errorReader) Read(p []byte) (n int, err error) {
	return 0, io.ErrUnexpectedEOF
}