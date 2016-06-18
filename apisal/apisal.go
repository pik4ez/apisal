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
	Url   string `json:"url"`
	W     int    `json:"w"`
	H     int    `json:"h"`
	Title string `json:"title"`
}

// Object represents object.
type Object struct {
	Point       Point   `json:"point"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Images      []Image `json:"images"`
}

// WikiObject
type WikiObject struct {
	Title string `json:"title"`
}
