# Parse

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

## Usage

    please parse [-i type] [-o type] [path...]

    -i type  Parse the input as 'type' (default: auto)
    -o type  Use 'type' as the output format (default: bash)

    If path is given, only output data from the path downwards.

## Examples

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
