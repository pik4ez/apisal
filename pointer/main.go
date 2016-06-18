package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"time"

	"github.com/pik4ez/apisal/apisal"
	gpx "github.com/ptrv/go-gpx"
)

func createPoint(waypoint gpx.Wpt) apisal.Point {
	return apisal.Point{
		Lat: waypoint.Lat,
		Lon: waypoint.Lon,
		// todo: use waypoint.Timestamp,
		Time: time.Now(),
	}
}

func main() {
	gpxFilePath := flag.String("gpx-file", "", "GPX file path")
	flag.Parse()

	if *gpxFilePath == "" {
		flag.Usage()
		// fmt.Println("Use --gpx-file parameter to specify gpx file path")
		os.Exit(1)
	}

	route, err := gpx.ParseFile(*gpxFilePath)
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
