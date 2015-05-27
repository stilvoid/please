#!/bin/bash

function dump {
    name=${1:?Must supply a variable name}

    declare -p $name | sed -e "s/^declare -. $name *= *//" | sed -e "s/^'\(.*\)'$/\\1/"
}

function load {
    obj="${1:?Must supply an object}"
    key="${2:?Must supply a key}"

    if [[ "$key" =~ "." ]]; then
        head=${key%%.*}
        rest=${key#*.}
    else
        head="$key"
        rest=""
    fi

    declare -A new=$obj

    if [ -n "$rest" ]; then
        load "${new[$head]}" "$rest"
    else
        echo "${new[$head]}"
    fi
}

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

read -r in

print "$in"
