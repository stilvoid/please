package internal

import (
	"errors"
	"io"
	"os"

	"github.com/andrew-d/go-termutil"
)

func StdinOrNothing() (io.Reader, error) {
	if termutil.Isatty(os.Stdin.Fd()) {
		return nil, nil
	}

	return os.Stdin, nil
}

func FileOrStdin(args ...string) (io.Reader, error) {
	if len(args) == 0 || args[0] == "-" {
		if termutil.Isatty(os.Stdin.Fd()) {
			return nil, errors.New("Unable to read from stdin")
		}

		return os.Stdin, nil
	}

	return os.Open(args[0])
}

func ReadFileOrStdin(args ...string) ([]byte, error) {
	reader, err := FileOrStdin(args...)
	if err != nil {
		return []byte{}, err
	}

	return io.ReadAll(reader)
}
