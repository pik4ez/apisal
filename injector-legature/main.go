package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"

	"github.com/pik4ez/apisal/apisal"
)

// GeoCoder contains client to Google Maps API.
type GeoCoder struct {
	client *maps.Client
}

func main() {
	googleMapsAPIKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if googleMapsAPIKey == "" {
		log.Fatal("GOOGLE_MAPS_API_KEY environment variable must be set")
	}
	geocoder, err := NewGeoCoder(googleMapsAPIKey)
	if err != nil {
		log.Fatal(err)
	}
	if s, err := os.Stdin.Stat(); err != nil || (s.Mode()&os.ModeCharDevice) != 0 {
		log.Fatal("stdin is empty!")
	}

	objectsReader := apisal.NewObjectsReader(os.Stdin)
	objectsWriter := apisal.NewObjectsWriter(os.Stdout)

	iteration := 0
	for {
		object, err := objectsReader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if iteration == 0 {
			ungeo, err := geocoder.UnGeocode(object.Point)
			// Silently omit legature if error.
			if err == nil {
				streetName := getStreetName(ungeo)
				if err == nil {
					legatureDescr := fmt.Sprintf("Начало пути: %s", streetName)
					legatureObject := apisal.Object{
						Type:        apisal.ObjectTypeLegature,
						Description: legatureDescr,
					}
					objectsWriter.WriteObject(legatureObject)
				}
			}
		}
		objectsWriter.WriteObject(object)
		iteration++
	}
}

// NewGeoCoder returns GeoCoder.
func NewGeoCoder(key string) (GeoCoder, error) {
	client, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		return GeoCoder{}, err
	}
	return GeoCoder{client: client}, nil
}

// UnGeocode returns geo objects by coordinates.
func (geocoder GeoCoder) UnGeocode(point apisal.Point) ([]maps.GeocodingResult, error) {
	req := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: point.Lat,
			Lng: point.Lon,
		},
		ResultType: []string{
			"street_address",
			"route",
		},
		Language: "ru",
	}
	resp, err := geocoder.client.Geocode(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getStreetName(res []maps.GeocodingResult) string {
	for _, component := range res {
		for _, addrComponent := range component.AddressComponents {
			if addrComponent.Types[0] == "route" {
				return addrComponent.LongName
			}
		}
	}
	return ""
}
