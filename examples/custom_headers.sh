#!/bin/bash

please respond -muh 200 localhost:8000 &

sleep 0.5s

(
    echo Content-Type: application/json
    echo X-Please-Request: true
    echo User-Agent: Please
    echo
    cat example.json
) | please post -i http://localhost:8000/headers?true >/dev/null
