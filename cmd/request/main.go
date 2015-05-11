package main

import (
    "bufio"
    "code.google.com/p/getopt"
    "fmt"
    "github.com/andrew-d/go-termutil"
    "io"
    "io/ioutil"
    "net/http"
    "net/textproto"
    "os"
    "path"
    "strings"
)

func createRequest(method string, url string, input io.Reader, headers_included bool) *http.Request {
    var req *http.Request
    var headers map[string][]string
    var err error

    method = strings.ToUpper(method)

    if headers_included {
        if input == nil {
            fmt.Println("Error reading headers")
            os.Exit(1)
        }

        // Parse headers from input
        reader := textproto.NewReader(bufio.NewReader(input))
        headers, err = reader.ReadMIMEHeader()

        if err != nil {
            fmt.Println("Error parsing headers:", err)
            os.Exit(1)
        }
    }

    req, err = http.NewRequest(method, url, input)

    if err != nil {
        fmt.Println("Error creating request:", err)
        os.Exit(1)
    }

    if headers_included {
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

func printResponse(resp *http.Response, include_headers bool) {
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    if include_headers {
        resp.Header.Write(os.Stdout)
        fmt.Println()
    }

    fmt.Println(string(body))
}

func main() {
    // Flags
    include_headers := getopt.Bool('i', "", "Include headers in output")
    headers_included := getopt.Bool('h', "", "Headers are included in the input")

    // Cheat because it's better than writing *another* arg parser
    getopt.SetParameters("<url>")
    getopt.SetProgram(fmt.Sprintf("%s <method>", path.Base(os.Args[0])))

    opts := getopt.CommandLine

    // Get the command
    opts.Parse(os.Args)
    if opts.NArgs() < 1 {
        getopt.Usage()
        os.Exit(1)
    }
    method := opts.Arg(0)

    // Deal with flags and get the url
    opts.Parse(opts.Args())
    if opts.NArgs() < 1 {
        getopt.Usage()
        os.Exit(1)
    }
    url := opts.Arg(0)

    var req *http.Request
    if termutil.Isatty(os.Stdin.Fd()) {
        req = createRequest(method, url, nil, *headers_included)
    } else {
        req = createRequest(method, url, os.Stdin, *headers_included)
    }

    resp := getResponse(req)

    printResponse(resp, *include_headers)
}
