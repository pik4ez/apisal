package apisal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

const (
	// ObjectTypeOrganic is a type of objects parsed from external sources.
	ObjectTypeOrganic = iota
	// ObjectTypeLegature is a type of objects added by injector to make
	// connections between organic objects.
	ObjectTypeLegature
	// ObjectTypeUser is a type of objects added by user and very similar
	// to organic objects by format.
	ObjectTypeUser
)

// Point represents point.
type Point struct {
	Lat  float64   `json:"lat"`
	Lon  float64   `json:"lon"`
	Time time.Time `json:"time"`
}

// Image represents image.
type Image struct {
	Url   string `json:"url"`
	W     int    `json:"w"`
	H     int    `json:"h"`
	Title string `json:"title"`
}

// Object represents object.
type Object struct {
	Type        int8
	Point       Point   `json:"point"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Images      []Image `json:"images"`
}

// ReadPoints returns a list of points from stdin.
func ReadPoints(r io.Reader) ([]Point, error) {
	var points []Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		point := Point{}
		err := json.Unmarshal([]byte(scanner.Text()), &point)
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}
	return points, nil
}

// ReadObjects reads all objects from stdin.
func ReadObjects(r io.Reader) ([]Object, error) {
	var objects []Object
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		object := Object{}
		err := json.Unmarshal([]byte(scanner.Text()), &object)
		if err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return objects, nil
}

// WriteObjects writes provided objects to stdout.
func WriteObjects(objects []Object) {
	for _, object := range objects {
		str, err := json.Marshal(object)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(str))
	}
}
