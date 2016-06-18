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
	avgPoint, counter := apisal.Point{}, 0
	for _, object := range objs {
		avgPoint.Lat += object.Point.Lat
		avgPoint.Lon += object.Point.Lon
		counter++
	}
	avgPoint.Lat = avgPoint.Lat / float64(counter)
	avgPoint.Lon = avgPoint.Lon / float64(counter)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.ParseFiles(filepath.Join(cwd, "./renderer-html/templates/simple.html")))
	data := struct {
		AvgPoint apisal.Point
		Objects  []apisal.Object
	}{
		avgPoint,
		objs,
	}
	t.Execute(os.Stdout, data)
}
