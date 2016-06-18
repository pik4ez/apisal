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

type PointsReader struct {
	scanner *bufio.Scanner
}
type PointsWriter struct {
	writer  io.Writer
}

func NewPointsReader(r io.Reader) *PointsReader {
	return &PointsReader{
		scanner: bufio.NewScanner(os.Stdin),
	}
}
func NewPointsWriter(w io.Writer) *PointsWriter {
	return &PointsWriter{
		writer: w,
	}
}

func (rw *PointsReader) ReadNext() (Point, error) {
	scan := rw.scanner.Scan()
	if scan {
		point := Point{}
		err := json.Unmarshal([]byte(rw.scanner.Bytes()), &point)
		if err != nil {
			return Point{}, io.EOF
		}
		return point, nil
	}
	return Point{}, io.EOF
}

func (rw *PointsWriter) Write(p Point) (error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return err
	}
	if _, err := rw.writer.Write(bytes); err != nil {
		return err
	}
	if _, err := rw.writer.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}

type ObjectsWriter struct {
	writer io.Writer
}

func NewObjectsWriter(w io.Writer) ObjectsWriter {
	return ObjectsWriter{
		writer: w,
	}
}

func (ow *ObjectsWriter) WriteObject(o Object) (error) {
	bytes, err := json.Marshal(o)
	if err != nil {
		return err
	}
	if _, err := ow.writer.Write(bytes); err != nil {
		return err
	}
	if _, err := ow.writer.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}


type ObjectsReader struct {
	scanner *bufio.Scanner
}

func NewObjectsReader(w io.Writer) ObjectsReader {
	return &ObjectsReader{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (r *ObjectsReader) ReadNext() (Object, error) {
	if r.scanner.Scan() {
		object := Object{}
		err := json.Unmarshal([]byte(r.scanner.Bytes()), &object)
		if err != nil {
			return Object{}, io.EOF
		}
		return object, nil
	}
	return Object{}, io.EOF
}





// ReadPoints returns a list of points from stdin.
// deprecated
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

// WritePoints writes provided points to stdout.
func WritePoints(points []Point) {
	for _, point := range points {
		jsonBytes, err := json.Marshal(point)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(jsonBytes))
	}
}

// ReadObjects reads all objects from stdin.
// deprecated
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
