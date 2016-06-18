package main

import (
	"log"
	"os"

	lib "github.com/pik4ez/apisal/apisal"
	"github.com/pik4ez/apisal/parser-wikimapia/mapia"
)

// APIKey contains a key to wikimapia API.
const APIKey = "59F5F0FD-B38A4635-6BC3D0EF-307471CE-7246D42E-A54DC82A-BB2A27C6-9A0FD0BE"
var usedWikiObjects = make(map[string]bool);

func main() {
	if s, err := os.Stdin.Stat(); err != nil || !(s.Mode() & os.ModeCharDevice) {
		log.Fatal("stdin is empty!")
	}

	points, err := lib.ReadPoints(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	objects := make([]lib.Object, 0, 100)
	for _, point := range points {
		if pointObjects, err := PointObjects(point); err == nil {
			objects = append(objects, pointObjects...)
			// fmt.Println(pointObjects)
		} else {
			log.Fatal(err)
		}
	}

	lib.WriteObjects(objects)
}

// PointObjects returns a list of objects by points from wikimapia.
func PointObjects(point lib.Point) ([]lib.Object, error) {
	var objects []lib.Object

	m := mapia.NewMapia(APIKey)
	places, err := m.GetNearbyObjects(point.Lat, point.Lon, "ru")
	if err != nil {
		return nil, err
	}

	for _, place := range places.Places {
		if _, ok := usedWikiObjects[place.ID]; ok {
			continue
		}
		usedWikiObjects[place.ID] = true

		extened, err := m.GetPlaceById(place.ID, "ru")
		if err != nil {
			return nil, err
		}
		var images []lib.Image
		if len(extened.Photos) > 0 {
			for _, photo := range extened.Photos {
				images = append(images, lib.Image{Url: photo.BigURL, H: 0, W: 0})
			}
		}
		o := lib.Object{Point: point, Title: place.Title, Description: extened.Description, Images: images}

		objects = append(objects, o)
	}

	return objects, nil
}
