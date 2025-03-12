# Please

Please is a versatile command-line utility for working with web requests and common data formats. It simplifies the process of making HTTP requests, serving content, and converting between different data formats.

## Features

- Make HTTP requests with support for different methods (GET, POST, PUT, DELETE)
- Start a simple HTTP server to serve local files
- Listen and respond to HTTP requests
- Parse and convert between various data formats
- Identify data format of files or input streams
- Query data using JMESPath expressions

## Installation

Choose one of the following methods:

### Using Homebrew

```bash
brew install stilvoid/tools/please
```

### Using Go

```bash
go install github.com/stilvoid/please@latest
```

### Manual Installation

Download the appropriate binary for your platform from the [releases page](https://github.com/stilvoid/please/releases).

## Commands

### please identify

Identifies the format of structured data from a file or stdin.

```bash
please identify [FILENAME]
```

### please parse

Parses and converts structured data between different formats.

```bash
please parse [FILENAME] [flags]
```

Supported formats:

Input formats:
- CSV
- HTML
- JSON
- MIME
- Query string
- TOML
- XML
- YAML

Output formats:
- Bash
- DOT (graph description)
- JSON
- Query string
- TOML
- XML
- YAML

Flags:
- `-f, --from string`: Input format (default "auto")
- `-t, --to string`: Output format (default "auto")
- `-q, --query string`: Apply a JMESPath query to the data

### please request

Sends HTTP requests and displays the response.

```bash
please request [method] [url] [flags]
```

Aliases: `get`, `post`, `put`, `delete`

Flags:
- `-b, --body string`: File to read request body from (use "-" for stdin)
- `-i, --include-headers`: Read headers from the request body
- `-v, --verbose`: Show response status line and headers

### please respond

Creates a temporary HTTP server that listens for requests and responds with specified content.

```bash
please respond [flags]
```

Flags:
- `-a, --address string`: Address to listen on
- `-p, --port int`: Port to listen on (default 8000)
- `-b, --body string`: File to read response body from (use "-" for stdin)
- `-i, --include-headers`: Read headers from the response body
- `-s, --status int`: HTTP status code to respond with (default 200)
- `-v, --verbose`: Show request headers

### please serve

Starts a simple HTTP server to serve local files.

```bash
please serve [PATH] [flags]
```

Flags:
- `-a, --address string`: Address to listen on
- `-p, --port int`: Port to listen on (default 8000)

## Examples

1. Convert JSON to YAML:
```bash
echo '{"name": "test"}' | please parse --from json --to yaml
```
Output:
```yaml
name: test
```

2. Query JSON data using JMESPath:
```bash
echo '{"users": [{"name": "test1"}, {"name": "test2"}]}' | please parse --query "users[*].name"
```
Output:
```json
[
  "test1",
  "test2"
]
```

3. Make a POST request with JSON body:
```bash
echo '{"data": "example"}' | please request post https://api.example.com/endpoint
```

4. Serve current directory:
```bash
please serve
```
This will start a web server on port 8000 serving the current directory.

5. Start a mock server that responds with specific content:
```bash
echo 'Hello World' | please respond -p 8080
```
This will start a server on port 8080 that responds with "Hello World" to all requests.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms found in the [LICENSE](LICENSE) file.