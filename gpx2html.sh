#!/usr/bin/env bash

./pointer/pointer -gpx-file $1 \
    | ./smoother/smoother -min-distance 100 \
    > ./cache/points_for_display.txt

./pointer/pointer -gpx-file $1 \
    | ./smoother/smoother -min-distance 500 \
	| ./parser-wikimapia/parser-wikimapia \
	| ./injector-legature/injector-legature \
	| ./filter/filter > ./cache/objects_prod.txt

cat ./cache/objects_prod.txt | ./renderer-babe/renderer-babe.php > $2

./renderer-html/renderer-html \
	-o ./cache/objects_prod.txt \
	-p ./cache/points_for_display.txt \
	-t ./renderer-html/templates/simple.html
