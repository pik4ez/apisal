package main

import (
	"fmt"
	"log"
	gpx "github.com/ptrv/go-gpx"
	"time"
)

// {"lat":0.5, "lon": 0.5, "time": "2016-06-17T20:30:25+0000"}
type Point struct {
	Latitude float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Time time.Time `json:"time"`
}

func main() {
	route, err := gpx.ParseFile("./simple.gpx")
	if err != nil {
		log.Fatal(err)
	}

	for _, track := range route.Tracks {
		for _, segment := range track.Segments {
			for _, waypoint := range segment.Waypoints {
				fmt.Println(waypoint.Lat)
				fmt.Println(waypoint.Lon)
				fmt.Println(waypoint.Timestamp)
			}
		}
	}
}