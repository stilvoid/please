#!/bin/bash

cat example.json | please respond 200 localhost:8000 &

sleep 0.5s

please get -sh http://localhost:8000
