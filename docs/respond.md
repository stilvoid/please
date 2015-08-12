# Respond

`please respond` sets up a web server that looks out for any one web request on the address and port specified, outputs the request that was received, returns the specified response, and then shuts down. This can be very useful for testing.

## Usage

    please respond [options...] <status> [<address>[:<port>]]

    -i    Headers included in the input

    -m    Include request method in output
    -u    Include URL in output
    -h    Include headers in output

## Examples

A "Hello, world" service:

Note: by default, `please respond` listens on `0.0.0.0:8000`

    $ echo "Hello, world" | please respond 200

Then in another terminal:

    $ please get http://127.0.0.1:8000/
    Hello, world
