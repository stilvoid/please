# ToDo

* Allow * in path filtering, e.g.:

    Given the data:

        [
            {
                "id": 1
            },
            {
                "id": 2
            }
        ]

    `please parse *.id` would result in:

        [1, 2]

* Tidy up help text - separate input/output params

    Use [optional] and <value goes here> conventions

    Replace [parameters...] in parse with something a bit more descriptive

* Split docs up and use readthedocs

* Write some tests!

* Fix stdin/stdout usage

* Make indenting optional across all formatters (but not bash I guess)
