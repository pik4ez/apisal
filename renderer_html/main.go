package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
)

// Point represents point.
type Point struct {
	Lat  float64
	Lon  float64
	Time string
}

// Image represents image.
type Image struct {
	Url   string
	W     int
	H     int
	Title string
}

// Object represents object.
type Object struct {
	Point       Point
	Title       string
	Description string
	Images      []Image
}

func main() {
	var inputObjOne = []byte("{\"point\": {\"lat\": 10.0, \"lng\": 10.0, \"time\": \"2016-06-17T20:30:25+0000\"}, \"title\": \"obj_one_title\", \"description\": \"obj_one_description\", \"images\": [{\"w\": 100, \"h\": 50, \"title\": \"img_one_title\", \"url\": \"http://example.org/image.png\"}]}")
	var objs []Object
	var objOne Object
	json.Unmarshal(inputObjOne, &objOne)
	objs = append(objs, objOne)
	fmt.Println(objOne.Images[0])
	t := template.Must(template.ParseFiles("templates/simple.html"))
	t.Execute(os.Stdout, objs)
}
