#!/usr/bin/env bash

go run pointer/main.go -gpx-file $1 | go run parser-wikimapia/main.go | go run renderer-html
