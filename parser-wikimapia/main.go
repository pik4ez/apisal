package main

import (
	"io"
	"log"
	"math/rand"
	"os"

	"time"

	lib "github.com/pik4ez/apisal/apisal"
	"github.com/pik4ez/apisal/parser-wikimapia/mapia"
)

var usedWikiObjects = make(map[int]bool)

// APIKeys contain a keys to wikimapia API.
var APIKeys = [...]string{
	"4599919F-BCFC13DC-577E3FA9-80FBD3BF-31DD56BF-D49033F8-AA3E31C5-4CD23144",
	"59F5F0FD-B38A4635-6BC3D0EF-307471CE-7246D42E-A54DC82A-BB2A27C6-9A0FD0BE",
	"F99FD50C-3926D677-8DF5BC36-0C3E1FFE-1A740EE5-D8B8C3F6-B09FEC2C-F2068243",
}

func getRandomAPIKey() string {
	rand.Seed(time.Now().UTC().UnixNano())
	return APIKeys[rand.Intn(len(APIKeys))]
}

func main() {
	apiKey := getRandomAPIKey()

	if s, err := os.Stdin.Stat(); err != nil || (s.Mode()&os.ModeCharDevice) != 0 {
		log.Fatal("stdin is empty!")
	}

	pointsReader := lib.NewPointsReader(os.Stdin)
	objectsWriter := lib.NewObjectsWriter(os.Stdout)

	for {
		point, err := pointsReader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		pointObjects, err := PointObjects(point, apiKey)
		if err != nil {
			log.Fatal(err)
		}
		for _, object := range pointObjects {
			objectsWriter.WriteObject(object)
		}
	}
}

// PointObjects returns a list of objects by points from wikimapia.
func PointObjects(point lib.Point, apiKey string) ([]lib.Object, error) {
	var objects []lib.Object

	m := mapia.NewMapia(apiKey)
	places, err := m.GetNearbyObjects(point.Lat, point.Lon, 30, 1, "ru")
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
		//fmt.Printf("%v", extened)

		var images []lib.Image
		if len(extened.Photos) > 0 {
			for _, photo := range extened.Photos {
				images = append(images, lib.Image{URL: photo.BigURL, H: 0, W: 0})
			}
		}
		o := lib.Object{
			Type:        lib.ObjectTypeOrganic,
			Point:       point,
			Lat:         place.Location.Lat,
			Lon:         place.Location.Lon,
			Title:       place.Title,
			Description: extened.Description,
			Images:      images,
		}

		objects = append(objects, o)
	}

	return objects, nil
}
