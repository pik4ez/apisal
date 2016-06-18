package apisal

import "time"

// Point represents point.
type Point struct {
	Lat  float64   `json:"lat"`
	Lon  float64   `json:"lon"`
	Time time.Time `json:"time"`
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

// WikiObject
type WikiObject struct {
	Title string `json:"title"`
}
