#!/bin/bash

please respond -muh 200 localhost:8000 &

sleep 0.5s

cat example.json | please post http://localhost:8000/api >/dev/null
