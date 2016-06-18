#!/usr/bin/env bash

./pointer/pointer -gpx-file $1 | ./parser-wikimapia/parser-wikimapia | ./renderer-html/renderer-html

