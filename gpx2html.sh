#!/usr/bin/env bash

./pointer/pointer -gpx-file $1 \
	| ./parser-wikimapia/parser-wikimapia \
	| ./injector-legature/injector-legature \
	| ./filter/filter > ./cache/objects_prod.txt

./renderer-html/renderer-html \
	-o ./cache/objects_prod.txt \
	-p ./cache/points.txt \
	-t ./renderer-html/templates/simple.html

