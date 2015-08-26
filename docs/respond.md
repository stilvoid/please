# Respond

`please respond` sets up a web server that looks out for any one web request on the address and port specified, outputs the request that was received, returns the specified response, and then shuts down. This can be very useful for testing.

## Usage

    please respond [option...] <STATUS> [<ADDRESS>[:<PORT>]]

    Listens on the specified address and port and responds with the chosen status code.
    Any data on stdin will be used as the body of the response.
    The request body will be printed to stdout.

    Input options:
        -i    Include headers from input

    Output options:
        -m    Output the request method
        -u    Output the requested path
        -h    Output headers with the request

## Examples

A "Hello, world" service:

Note: by default, `please respond` listens on `0.0.0.0:8000`

    $ echo "Hello, world" | please respond 200

Then in another terminal:

    $ please get http://127.0.0.1:8000/
    Hello, world
