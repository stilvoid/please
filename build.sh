#!/bin/bash

# This script will build please for all platforms

platforms="linux darwin windows"
architectures="386 amd64"

echo "Building please"

for platform in $platforms; do
    for architecture in $architectures; do
        echo "... $platform $architecture..."

        name=please-$platform-$architecture

        if [ "$platform" == "windows" ]; then
            name=${name}.exe
        fi

        GOOS=$platform GOARCH=$architecture go build -o $name
    done
done

echo "All done."
