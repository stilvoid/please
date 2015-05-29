#!/bin/bash

function print {
    local obj="${1:?Must supply an object}"
    local indent="${2-}"

    declare -A new=$obj

    if [[ "${#new[@]}" == "1" && "${!new[@]}" == "0" ]]; then
        # It's (probably) a value
        echo "$indent${new[0]}"
    else
        # It's an object
        for key in "${!new[@]}"; do
            echo "$indent$key:"
            print "${new[$key]}" "$indent  "
        done
    fi
}

in=$(cat example.json | please parse -i json -o bash)

print "$in"
