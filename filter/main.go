package main

import (
	"flag"
	"io"
	"log"
	"os"
	"unicode/utf8"

	lib "github.com/pik4ez/apisal/apisal"
)

func main() {
	var photosMin int
	var descrLenMin int
	var titleLenMin int
	flag.IntVar(&photosMin, "p", 0, "Minimum photos per object")
	flag.IntVar(&titleLenMin, "t", 10, "Minimum title length")
	flag.IntVar(&descrLenMin, "d", 128, "Minimum description length")
	flag.Parse()

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

		titleLen := utf8.RuneCountInString(object.Title) > 0
		//descLen := utf8.RuneCountInString(object.Description) > 0
		imagesCount := len(object.Images)
		imagesExist := imagesCount > 0

		if imagesExist && titleLen {
			writer.WriteObject(object)
			continue
		}

		if utf8.RuneCountInString(object.Title) < titleLenMin {
			continue
		}
		if utf8.RuneCountInString(object.Description) < descrLenMin {
			continue
		}

		writer.WriteObject(object)
	}
}
