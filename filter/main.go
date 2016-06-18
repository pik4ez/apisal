package main

import (
	lib "github.com/pik4ez/apisal/apisal"
	"os"
	"io"
	"log"
	"unicode/utf8"
)

func main() {
	reader := lib.NewObjectsReader(os.Stdin)
	writer := lib.NewObjectsWriter(os.Stdout)
	for {
		object, err := reader.ReadNext();
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		title := len(object.Title) > 0
		desc := len(object.Description) > 0
		images := len(object.Images) > 0

		//fmt.Println(title, desc, images)

		if !title {
			continue
		}

		if !desc && !images {
			continue
		}

		if utf8.RuneCountInString(object.Title) < 10 && !desc && !images {
			continue
		}

		writer.WriteObject(object)
	}
}
