package main

import (
	"html/template"
	"log"
	"os"
	lib "github.com/pik4ez/apisal/apisal"
	"flag"
	"io"
)

func main() {
	pFilename := flag.String("p", "", "Points filename")
	oFilename := flag.String("o", "", "Objects filename")
	tFilename := flag.String("t", "", "Template filename")
	flag.Parse()

	if *tFilename == "" || *oFilename == "" || *pFilename == "" {
		flag.Usage()
		os.Exit(1)
	}

	objects, err := loadObjects(*oFilename)
	points, err := loadPoints(*pFilename)

	context := struct {
		Points  []lib.Point
		Objects []lib.Object
	}{
		points,
		objects,
	}

	template := template.Must(template.ParseFiles(*tFilename))
	err = template.Execute(os.Stdout, context)
	if err != nil {
		log.Fatal(err)
	}
}

func loadPoints(filename string) ([]lib.Point, error) {
	points := make([]lib.Point, 0, 100)
	pFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer pFile.Close()
	pReader := lib.NewPointsReader(pFile)
	for {
		point, err := pReader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		points = append(points, point)
	}

	return points, nil
}

func loadObjects(filename string) ([]lib.Object, error) {
	objects := make([]lib.Object, 0, 100)
	oFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer oFile.Close()
	oReader := lib.NewObjectsReader(oFile)
	for {
		object, err := oReader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		objects = append(objects, object)
	}

	return objects, nil
}
