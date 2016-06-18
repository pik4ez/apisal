package main

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

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
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(objs); err != nil {
		fmt.Printf("error: %v\n", err)
	}
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
