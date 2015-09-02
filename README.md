# Please

Please is a command line utility that makes it easy to integrate web APIs into your shell scripts.

It's called Please because the web works much better if you ask nicely.

It is comprised of three sub-commands:

* `please request`

    for communicating with web servers

* `please respond`

    acts as a one-shot web server - useful in testing you applications

* `please identify`

    given structured data on standard input, this will output the format used (e.g. "json", "yaml")

* `please parse`

    understands the data exchange formats of the web and can translate between them

## Installing

**Go users**: Simply `go install github.com/stilvoid/please`.

**Arch Linux**: There's a [please package in the AUR](https://aur4.archlinux.org/packages/please/).

**Anyone else**: Grab the appropriate download from the [latest release](https://github.com/stilvoid/please/releases) and put it somewhere in your path.

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
