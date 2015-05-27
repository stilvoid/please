# please-respond

## Usage

    please-respond [options...] <status> [<address>[:<port>]]

    please-respond will wait for an HTTP request on <address>:<port> and
    will output the contents of the request.

    If there is any data on standard in (e.g. please-respond is on the
    right-hand side of a pipe), that will be sent as the response body.

    <status> should be a 3 digit number and will be used as the response code.

    Standard statuses include: 200, 404, 301, 302, 500, etc.

    <address> is the address to listen on. It defaults to 0.0.0.0

    <port> is the port to listen on. It default to 8080

    Options:

        -i      The data on standard input contains headers and they should be
                included in the response - overriding default headers where
                appropriate.

        -h      Output headers from the request.

        -m      Output the method of the request

        -u      Output the URL of the request
