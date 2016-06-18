package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gotokatsuya/gosquare/dispatcher"
	"github.com/gotokatsuya/gosquare/service/venues"
	"github.com/pik4ez/apisal/apisal"
)

func main() {
	if s, err := os.Stdin.Stat(); err != nil || s.Size() == 0 {
		log.Fatal("stdin is empty!")
	}

	points, err := apisal.ReadPoints(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	objects := make([]apisal.Object, 0, 100)
	for _, point := range points {
		if pointObjects, err := VenuesExplore(point); err == nil {
			objects = append(objects, pointObjects...)
		} else {
			log.Fatal(err)
		}
	}

	apisal.WriteObjects(objects)
}

// VenuesExplore returns venues near the specified location.
func VenuesExplore(point apisal.Point) ([]apisal.Object, error) {
	var objects []apisal.Object
	client := dispatcher.NewClient()
	req := venues.NewExploreRequest()
	req.LatLng = fmt.Sprintf("%3f,%3f", point.Lat, point.Lon)
	res, err := venues.Explore(client, req)
	if err != nil {
		return nil, err
	}
	for _, v := range res.GetVenues() {
		fmt.Println(v.Tips)
		o := apisal.Object{Point: point, Title: v.Name}
		objects = append(objects, o)
	}
	return objects, nil
}
