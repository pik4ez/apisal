.PHONY: all pointer parser-wikimapia injector-legature renderer-html www points objects pipeline clean filter smoother

all: pointer parser-wikimapia injector-legature filter renderer-html www filter smoother

pointer:
	go build -o ./pointer/pointer ./pointer/

parser-wikimapia:
	go build -o ./parser-wikimapia/parser-wikimapia ./parser-wikimapia/

injector-legature:
	go build -o ./injector-legature/injector-legature ./injector-legature/

filter:
	go build -o ./filter/filter ./filter/

smoother:
	go build -o ./smoother/smoother ./smoother/

renderer-html:
	go build -o ./renderer-html/renderer-html ./renderer-html/

www:
	go build -o ./www/www ./www/

points:
	go run ./pointer/main.go -gpx-file ./pointer/simple.gpx \
		| tee -p /dev/tty \
		| go run ./smoother/main.go -min-distance 200 \
		| tee /dev/tty > ./cache/points.txt

objects:
	cat ./cache/points.txt \
		| tee -p /dev/tty \
		| go run ./parser-wikimapia/main.go \
		| tee -p /dev/tty \
		| go run ./filter/main.go > ./cache/objects.txt

render:
	cat ./cache/objects.txt \
		| tee -p /dev/tty \
		| go run ./renderer-html/main.go -o ./cache/objects.txt -p ./cache/points.txt \
			-t ./renderer-html/templates/simple.html > ./html/render.html
	@echo
	@echo "file://"$(realpath ./html/render.html)

pipeline: points objects render

clean:
	rm -rf ./cache/*
