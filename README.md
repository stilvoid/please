# Please

Please is a command line utility that makes it easy to integrate web APIs into your shell scripts.

It's called Please because the web works much better if you ask nicely.

It is comprised of three sub-commands:

* `please request`

    for communicating with web servers

* `please respond`

    acts as a one-shot web server - useful in testing you applications

* `please parse`

    understands the data exchange formats of the web and can translate between them

## Examples

There are some examples in the examples folder.

Here are a few other ways that please might be useful in a bash script:

### Getting the title of a web page

    $ please get http://example.com/ | please parse html.head.title
    Example Domain

### Testing an api

    $ echo '{"thing": 1}' | please post http://myapi.com
    Success

### Testing authentication

    $ echo 'Authorization: Bearer mytoken' | please get -is http://myapi.com
    401

### Providing a single-use mock web server for testing client code

    $ (echo Hello, world | please respond 200) & sleep 1s && curl -i http://localhost:8000
    HTTP/1.1 200 OK
    Date: Fri, 29 May 2015 22:23:30 GMT
    Content-Length: 13
    Content-Type: text/plain; charset=utf-8

    Hello, world

### Converting between structured data formats

    $ echo '{"some": ["lovely", "json"], "now": "yaml"}' | please parse -o yaml
    some:
    - lovely
    - json
    now: yaml

## please request

Probably the most important thing you will need to do with a web service is to communicate with it. `please request` is fluent in HTTP and allows you to send any type of request along with any headers and content you need. It then outputs the response, optionally including the status code and headers.

`please request` supports all HTTP method types and if you ever need a non-standard one, you can specify it directly.

### Usage

    please request <method> [options...] <url>

    -i    Headers included in the input

    -s    Output HTTP status line with the response
    -h    Output headers with the response

    There are aliases for the most common methods:

        please get <url>
        please post <url>
        please put <url>
        please delete <url>

    For any other methods, you will need to specify the method directly:

        please request head <url>
        please request info <url>
        please request patch <url>

### Examples

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

## `please respond`

`please respond` sets up a web server that looks out for any one web request on the address and port specified, outputs the request that was received, returns the specified response, and then shuts down. This can be very useful for testing.

### Usage

    please respond [options...] <status> [<address>[:<port>]]

    -i    Headers included in the input

    -m    Include request method in output
    -u    Include URL in output
    -h    Include headers in output

### Examples

A "Hello, world" service:

Note: by default, `please respond` listens on `0.0.0.0:8000`

    $ echo "Hello, world" | please respond 200

Then in another terminal:

    $ please get http://127.0.0.1:8000/
    Hello, world

## `please parse`

`please parse` can deal with a number of structured data formats and make them easy to parse from bash

* JSON
* YAML
* XML
* CSV
* HTML
* MIME messages

Please parses data structures in the above formats and can output them as:

* bash `declare` syntax
* JSON
* YAML
* XML
* dot (graphviz)

If you're not familiar with associative arrays in bash or how `declare` works, it's worth reading the following:

* <http://www.gnu.org/software/bash/manual/html_node/Arrays.html>

* <http://www.tldp.org/LDP/abs/html/declareref.html>

### Usage

    please parse [-i type] [-o type] [path...]

    -i type  Parse the input as 'type' (default: auto)
    -o type  Use 'type' as the output format (default: bash)

    If path is given, only output data from the path downwards.

### Examples

Parsing a whole tree:

    $ echo '{"this":{"is":["some","json"],"that":"we"},"will":"parse"}' | please parse
    ([this]="([is]=\"([0]=some [1]=json)\" [that]=we)" [will]=parse)

Getting a single value:

    $ echo '{"this":{"is":["some","json"],"that":"we"},"will":"parse"}' | please parse this.that
    we

Specifying the input format:

Note the `-i` flag

    $ echo '<xml example="true"><child>one</child><child>two</child></xml>' | please parse -i xml
    ([xml]="([-example]=\"true\" [child]=\"([0]=\\\"one\\\" [1]=\\\"two\\\")\")")

Specifying the output format:

Note the `-o` flag

    $ echo '{"json": ["input", "here"], "yaml": "output"}' | please parse -o yaml
    'json': 
      - "input"
      - "here"
    'yaml': "output"

Making use of the bash-declare output format

    $ echo '{"json": ["array", "values"]}' | please parse json | (declare -A data=$(cat -); echo ${data[1]})
    values

Generating a graph from some json (you need graphviz installed)

    $ echo '{"vars": ["foo", "bar", "baz"], "cake": {"is_lie": true}}' | please parse -o dot | dot -Tpng > graph.png
