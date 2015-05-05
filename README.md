# Please

Please is a command line utility that makes it easy to integrate web APIs into your shell scripts.

## Getting data

Please only speaks HTTP but it understands a variety of socket types:

* http://api.example.com/

* https://api.example.com/

* unix:///var/run/my.sock

Please supports all HTTP method types and if you ever need a non-standard one, you can specify:

* GET
* POST
* PUT
* DELETE
* PATCH
* HEAD
* OPTIONS

Please can customise all parts of the HTTP request (but provides sensible defaults if you don't):

* Headers
* Query string parameters
* Request body

### Examples

Simple HTTP GET:

    please get http://api.example.com/

Posting some data:

    please post "Hello, world" to http://api.example.com/

or:

    cat "Hello, world" | please post to http://api.example.com/

The `to` is optional but we think it reads better :)

## Parsing data

Please can deal with a number of structured data formats and make them easy to parse from bash

* JSON
* YAML
* XML
* HTML

Please parses data structures out into a format that can easily be parsed by bash.

Alternatively, please can parse out just the particular value or values you're interested in.

### Examples

Parsing a whole tree:

    echo '{"this":{"is":["some","json"],"that":"we"},"will":"parse"}' | please parse

    ([this]="([is]=\"([0]=some [1]=json)\" [that]=we)" [will]=parse)

Getting a single value:

    echo '{"this":{"is":["some","json"],"that":"we"},"will":"parse"}' | please parse this.that

    we

## Formatting data

Please can 
