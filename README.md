Please is a utility for making and receiving web requests and parsing and reformatting the common data formats that are sent over them.

## Installing

`brew install stilvoid/tools/please`

_or_

Download a binary from the [releases](https://github.com/stilvoid/please/releases) page.

_or_

Run `go install github.com/stilvoid/please@latest`

## Usage

```
please [command]

Available Commands:
  help        Help about any command
  identify    Identify the format of some structured data from FILENAME or stdin if omitted
  parse       Parse and convert structured data from FILENAME or stdin if omitted
  request     Send a web request to a url and print the response
  respond     Listen for an HTTP request and respond to it
  serve       Serve the contents of PATH (current directory if omitted) through a simple web server
```

## Please Identify

Identify the format of some structured data from FILENAME or stdin if omitted

```
Usage:
  please identify (FILENAME) [flags]

Flags:
  -h, --help   help for identify
```

## Please parse

Parse and converted structured data from FILENAME or stdin if omitted.

```
Input formats:
  csv
  html
  json
  mime
  query
  toml
  xml
  yaml

Output formats:
  bash
  dot
  json
  query
  toml
  xml
  yaml

## Please parse

Usage:
  please parse (FILENAME) [flags]

Flags:
  -f, --from string    input format (see please help parse for formats) (default "auto")
  -h, --help           help for parse
  -q, --query string   JMESPath query
  -t, --to string      output format (see please help parse for formats) (default "auto")
```

## Please Request

Send a web request to a url and print the response

```
Usage:
  please request [method] [url] [flags]

Aliases:
  request, get, post, put, delete

Flags:
  -b, --body string       Filename to read the request body from. Use - or omit for stdin.
  -h, --help              help for request
  -i, --include-headers   Read headers from the request body
  -v, --verbose           Output response status line and headers
```

## Please Respond

Listen for an HTTP request and respond to it

```
Usage:
  please respond [flags]

Flags:
  -a, --address string    Address to listen on
  -b, --body string       Filename to read the response body from. Use - or omit for stdin
  -h, --help              help for respond
  -i, --include-headers   Read headers from the response body
  -p, --port int          Port to listen on (default 8000)
  -s, --status int        Status code to respond with (default 200)
  -v, --verbose           Output request headers
```

## Please Serve

Serve the contents of PATH (current directory if omitted) through a simple web server

```
Usage:
  please serve (PATH) [flags]

Flags:
  -a, --address string   Address to listen on
  -h, --help             help for serve
  -p, --port int         Post to listen on (default 8000)
```
