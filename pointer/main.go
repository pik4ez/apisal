package main

import (
	"flag"
	"log"
	"os"

	"time"

	"github.com/pik4ez/apisal/apisal"
	gpx "github.com/ptrv/go-gpx"
)

// NewPointByGpx creates point object by GPX point.
func NewPointByGpx(waypoint gpx.Wpt) apisal.Point {
	return apisal.Point{
		Lat: waypoint.Lat,
		Lon: waypoint.Lon,
		// TODO: use waypoint.Timestamp,
		Time: time.Now(),
	}
}

func main() {
	gpxFilePath := flag.String("gpx-file", "", "GPX file path")
	flag.Parse()

	if *gpxFilePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	route, err := gpx.ParseFile(*gpxFilePath)
	if err != nil {
		log.Fatal(err)
	}

	var points []apisal.Point
	for _, track := range route.Tracks {
		for _, segment := range track.Segments {
			for _, waypoint := range segment.Waypoints {
				point := NewPointByGpx(waypoint)
				points = append(points, point)
			}
		}
	}

	apisal.WritePoints(points)
}
