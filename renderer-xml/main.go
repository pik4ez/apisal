package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

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
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(objs); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
