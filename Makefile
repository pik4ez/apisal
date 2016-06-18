.PHONY: all pointer parser-wikimapia injector-legature filter renderer-html www

all: pointer parser-wikimapia injector-legature filter renderer-html www

pointer:
	go build -o ./pointer/pointer ./pointer/

parser-wikimapia:
	go build -o ./parser-wikimapia/parser-wikimapia ./parser-wikimapia/

injector-legature:
	go build -o ./injector-legature/injector-legature ./injector-legature/

filter:
	go build -o ./filter/filter ./filter/

renderer-html:
	go build -o ./renderer-html/renderer-html ./renderer-html/

www:
	go build -o ./www/www ./www/
