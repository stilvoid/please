# ToDo

* Add filtering to parse

        e.g. echo mystuff.json | please parse person.name=steve

        i.e. if there's an equals, it's a filter

        if there's no equals, it restricts output

        there can only be one output restriction

        there can be multiple filters

        filters are OR'ed

        filtering is case-insensitive

* Add regex filters

    e.g. echo mystuff.com | please parse person.name~^steve

    case-insensitive

* Add benchmarks for stuff in common

* Add an optional `FILE` parameter to everything

* Mention expectations of data on stdin in help text

* Write some more tests!

    * Particularly for the commands

    * Cover all formatters and parsers

* Fix stdin/stdout/stderr usage

    * Make sure all errors go to stderr

* Make indenting optional across all formatters (but not bash I guess)

* Add a preferences thing for storing headers etc?

    Could store commonly used headers for a domain, e.g. auth
