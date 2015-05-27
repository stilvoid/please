# please-request

## Usage

    please-request <method> [options...] <url>

    please-request will make an HTTP request to <url> and output the response.
    If there is any data on standard in (e.g. please-request is on the
    right-hand side of a pipe), that will be sent as the request body.

    <method> should be a single word and it will be converted to upper-case.

    Standard methods include: GET, POST, PUT, DELETE, etc.

    Options:

        -i      The data on standard input contains headers and they should be
                included in the request - overriding default headers where
                appropriate.

        -s      Include the HTTP status code with the response

        -h      Output headers with the response.

    Notes:

        No matter in which order the -s and -h options are given, output
        will always take the same format:

        status code (if option enabled)
        headers... (if option enabled)
        blank line (if headers option enabled)
        body
