.PHONY: all pointer parser-wikimapia renderer-html www

all: pointer parser-wikimapia renderer-html www

pointer:
	go build -o ./pointer/pointer ./pointer/

parser-wikimapia:
	go build -o ./parser-wikimapia/parser-wikimapia ./parser-wikimapia/

renderer-html:
	go build -o ./renderer-html/renderer-html ./renderer-html/

www:
	go build -o ./www/www ./www/
