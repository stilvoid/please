#!/bin/bash

# This script will build please for all platforms

declare -A platforms=([linux]=linux [darwin]=osx [windows]=windows)
declare -A architectures=([386]=i386 [amd64]=amd64)

echo "Building please"

for platform in ${!platforms[@]}; do
    for architecture in ${!architectures[@]}; do
        echo "... $platform $architecture..."

        name=please-${platforms[$platform]}-${architectures[$architecture]}

        if [ "$platform" == "windows" ]; then
            name=${name}.exe
        fi

        GOOS=$platform GOARCH=$architecture go build -o $name
    done
done

echo "All done."
