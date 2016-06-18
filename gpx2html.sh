#!/usr/bin/env bash

./pointer/pointer -gpx-file $1 | ./parser-wikimapia/parser-wikimapia | ./injector-legature/injector-legature | ./filter | ./renderer-html/renderer-html

