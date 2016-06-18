package main

import (
	"flag"
	"io"
	"log"
	"os"

	geo "github.com/kellydunn/golang-geo"
	"github.com/pik4ez/apisal/apisal"
)

func main() {
	minDistance := flag.Float64("min-distance", 0, "Min distance between pointers")
	flag.Parse()

	if *minDistance <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	reader := apisal.NewPointsReader(os.Stdin)
	writer := apisal.NewPointsWriter(os.Stdout)

	var prevGeoPoint *geo.Point

	for {
		point, err := reader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		currGeoPoint := geo.NewPoint(point.Lat, point.Lon)

		if prevGeoPoint != nil {
			distance := currGeoPoint.GreatCircleDistance(prevGeoPoint)
			if distance*1000 < *minDistance {
				continue
			}
		}

		writer.Write(point)

		prevGeoPoint = currGeoPoint
	}
}
