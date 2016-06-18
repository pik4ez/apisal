package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"github.com/pik4ez/apisal/apisal"
)

func main() {
	if os.Stdin == nil {
		log.Fatal("You should provide objects.")
	}
	objs, err := apisal.ReadObjects(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.ParseFiles(filepath.Join(cwd, "./renderer-html/templates/simple.html")))
	t.Execute(os.Stdout, objs)
}
