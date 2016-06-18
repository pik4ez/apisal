package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pik4ez/apisal/apisal"
)

func main() {
	if os.Stdin == nil {
		log.Fatal("You should provide objects.")
	}
	objs, err := ReadObjects(os.Stdin)
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

// ReadObjects reads all objects from stdin.
func ReadObjects(r io.Reader) ([]apisal.Object, error) {
	var objects []apisal.Object
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		object := apisal.Object{}
		err := json.Unmarshal([]byte(scanner.Text()), &object)
		if err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}
