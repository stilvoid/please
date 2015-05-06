package main

import (
    "flag"
    "fmt"
    "net/http"
)

func main() {
    // Deal with flags first
    headers := flag.String("h", "", "Command separated pairs of <header>=<value>")
    flag.Parse()

    // Remaining args
    method := flag.Arg(0)
    url := flag.Arg(1)

    fmt.Printf("%s %s\n", method, url)
    fmt.Println(*headers)

    // Create client
    client := &http.Client{}

    // Construct request - TODO: Add body
    req, _ := http.NewRequest(method, url, nil)

    // TODO - headers
    // req.Header.Add("Foo", "Bar")

    resp, _ := client.Do(req)

    fmt.Println(resp)
}
