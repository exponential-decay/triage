#!/usr/bin/env bash

java -jar tools/tika-server-1.13.jar --port=2041 > /dev/null 2>&1 &
sf -serve "localhost:2040" > /dev/null 2>&1 &
