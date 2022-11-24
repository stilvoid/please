package parse

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/textproto"
)

func Mime(input []byte) (interface{}, error) {
	inputReader := bufio.NewReader(bytes.NewReader(input))

	reader := textproto.NewReader(inputReader)

	headers, err := reader.ReadMIMEHeader()
	if err != nil {
		return nil, err
	}

	bytesBody, err := ioutil.ReadAll(inputReader)
	if err != nil {
		return nil, err
	}

	messageHeaders := make(map[string]interface{})
	for key, value := range headers {
		messageHeaders[key] = value
	}

	message := make(map[string]interface{})
	message["headers"] = messageHeaders
	message["body"] = string(bytesBody)

	return message, nil
}
