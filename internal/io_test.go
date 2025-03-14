package internal_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stilvoid/please/internal"
)

func TestReadFileOrStdin(t *testing.T) {
	// Create a temporary file with content
	content := []byte("test content")
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// Test reading from file
	data, err := internal.ReadFileOrStdin(tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !bytes.Equal(data, content) {
		t.Errorf("Expected %q, got %q", content, data)
	}

	// Test with non-existent file
	_, err = internal.ReadFileOrStdin("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

func TestFileOrStdin(t *testing.T) {
	// Create a temporary file with content
	content := []byte("test content")
	tmpFile, err := os.CreateTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	// Test reading from file
	reader, err := internal.FileOrStdin(tmpFile.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	data, err := io.ReadAll(reader)
	if err != nil {
		t.Errorf("Unexpected error reading from file: %v", err)
	}
	if !bytes.Equal(data, content) {
		t.Errorf("Expected %q, got %q", content, data)
	}

	// Test with non-existent file
	_, err = internal.FileOrStdin("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}

// Note: Testing StdinOrNothing and some aspects of FileOrStdin with actual stdin is challenging
// in automated tests because they check if stdin is a TTY. These functions would
// typically need to be tested with mocks or by temporarily redirecting stdin.