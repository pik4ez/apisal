package main

import (
	"fmt"
)

type Point struct {
	Lat  float64
	Lon  float64
	Time string
}

type Image struct {
	Url   string
	W     string
	H     string
	Title string
}

type Object struct {
	Point       Point
	Title       string
	Description string
	Images      []Image
}

func main() {
	fmt.Println("hillo")
}
