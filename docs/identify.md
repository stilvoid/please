# Identify

`please identify` can identify some common data formats and, give some data on standard input, will print out the format used.

Currently, `identify` can reliably tell the difference between:

* JSON
* YAML
* XML
* MIME messages

## Usage

    please identify

    Identifies the format of the structured data on standard input

## Examples

Identifying some data:

    $ echo '{"this":{"is":["some","json"],"that":"we"},"will":"parse"}' | please identify
    json

    $ echo -e "this:\n  looks: like\n  some: yaml" | please identify
    yaml
