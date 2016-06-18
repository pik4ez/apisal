package main

import (
	"io"
	"log"
	"os"
	"unicode/utf8"

	lib "github.com/pik4ez/apisal/apisal"
)

func main() {
	reader := lib.NewObjectsReader(os.Stdin)
	writer := lib.NewObjectsWriter(os.Stdout)
	for {
		object, err := reader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if object.Type != lib.ObjectTypeOrganic {
			writer.WriteObject(object)
			continue
		}

		title := utf8.RuneCountInString(object.Title) > 0
		desc := utf8.RuneCountInString(object.Description) > 0
		images := len(object.Images) > 0

		// fmt.Println(title, desc, images)

		if !title || !desc || !images {
			continue
		}

		if !title {
			continue
		}

		if !desc && !images {
			continue
		}

		if utf8.RuneCountInString(object.Title) < 10 && utf8.RuneCountInString(object.Description) < 128 && !images {
			continue
		}

		writer.WriteObject(object)
	}
}
