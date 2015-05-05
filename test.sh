#!/bin/bash

one_top="([0]=one [1]=top)"
ont_bottom="([0]=one [1]=bottom)"
one="([top]=\"$one_top\" [bottom]=\"$one_bottom\")"
two="([up]=\"two_up\" [down]=two_down)"
all="([one]=\"${one//\"/\\\"}\" [two]=\"${two//\"/\\\"}\")"

echo "One: $one"
echo "Two: $two"
echo "All: ${all}"

declare -A test=$(echo $all)

declare -p test

echo "Test: $test"

echo "Test[one]: ${test[one]}"

declare -A test=$(echo ${test[one]})

echo "Test2: $test"

echo "Test2[top]: ${test[top]}"

declare -A test=$(echo ${test[top]})

echo "Test3: $test"

echo "Test3[0]: ${test[0]}"
echo "Test3[1]: ${test[1]}"

echo "---"

function dump {
    name=${1:?Must supply a variable name}

    declare -p $name | sed -e "s/^declare -. $name *= *//" | sed -e "s/^'\(.*\)'$/\\1/"
}

function load {
    obj=${1:?Must supply an object}
    key=${2:?Must supply a key}

    if [[ "$key" =~ "." ]]; then
        head=${key%.*}
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

declare -a left=(one two three)
declare -a right=(four five six)
declare -A both=([left]=$(dump left) [right]=$(dump right))

foo=$(dump both)

echo "Dump: $foo"

load "$foo" left.0
load "$foo" left.1
