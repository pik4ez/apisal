package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	gpx "github.com/ptrv/go-gpx"
)

type Point struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Time      string  `json:"time"`
}

func createPoint(waypoint gpx.Wpt) Point {
	return Point{
		Latitude:  waypoint.Lat,
		Longitude: waypoint.Lon,
		Time:      waypoint.Timestamp,
	}
}

func main() {
	gpcFilePath := flag.String("gpx-file", "", "GPX file path")
	flag.Parse()

	if *gpcFilePath == "" {
		fmt.Println("Use --gpx-file parameter to specify gpx file path")
		os.Exit(1)
	}

	route, err := gpx.ParseFile(*gpcFilePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, track := range route.Tracks {
		for _, segment := range track.Segments {
			for _, waypoint := range segment.Waypoints {
				point := createPoint(waypoint)
				jsonBytes, err := json.Marshal(point)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(string(jsonBytes))
			}
		}
	}
}
