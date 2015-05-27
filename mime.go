package please

import (
    "bufio"
    "bytes"
    "io/ioutil"
    "net/textproto"
)

func ParseMIME(input []byte, path string) (interface{}, error) {
    input_reader := bufio.NewReader(bytes.NewReader(input))

    reader := textproto.NewReader(input_reader)

    headers, err := reader.ReadMIMEHeader()
    if err != nil {
        return nil, err
    }

    bytes_body, err := ioutil.ReadAll(input_reader)
    if err != nil {
        return nil, err
    }

    message_headers := make(map[string]interface{})
    for key, value := range headers {
        message_headers[key] = value
    }

    message := make(map[string]interface{})
    message["headers"] = message_headers
    message["body"] = string(bytes_body)

    return message, nil
}
