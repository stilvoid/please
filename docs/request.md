# Request

Probably the most important thing you will need to do with a web service is to communicate with it. `please request` is fluent in HTTP and allows you to send any type of request along with any headers and content you need. It then outputs the response, optionally including the status code and headers.

`please request` supports all HTTP method types and if you ever need a non-standard one, you can specify it directly.

## Usage

    please request <METHOD> [option...] <URL>

    Makes a web request to URL using METHOD

    Shortcut aliases:
        please get
        please post
        please put
        please delete

    Input options:
        -i    Include headers from input

    Output options:
        -s    Output HTTP status line
        -h    Output headers

## Examples

Simple HTTP GET:

    $ please get http://api.example.com/

Posting "Hello, world" to a URL:

    $ echo "Hello, world" | please post http://api.example.com/

Including headers as part of your request:

Note the `-i` flag.

    $ cat <<EOF | please post -i http://api.example.com/
    Content-Type: text/html

    Hello, world
    EOF
