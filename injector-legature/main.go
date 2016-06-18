package main

import (
	"io"
	"log"
	"os"
)

func main() {
	googleMapsAPIKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if googleMapsAPIKey == "" {
		log.Fatal("GOOGLE_MAPS_API_KEY environment variable must be set")
	}
	if s, err := os.Stdin.Stat(); err != nil || (s.Mode()&os.ModeCharDevice) != 0 {
		log.Fatal("stdin is empty!")
	}

	objectsReader := lib.NewObjectsReader(os.Stdin)
	objectsWriter := lib.NewObjectsWriter(os.Stdout)

	for {
		point, err := pointsReader.ReadNext()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		pointObjects, err := PointObjects(point)
		//pointObjects, err := mockObjects(point)
		if err != nil {
			log.Fatal(err)
		}
		for _, object := range pointObjects {
			objectsWriter.WriteObject(object)
		}
	}

	//c, err := maps.NewClient(maps.WithAPIKey(googleMapsAPIKey))
	//if err != nil {
	//log.Fatalf("fatal error: %s", err)
	//}
	//rOne := &maps.GeocodingRequest{
	//LatLng: &maps.LatLng{
	//Lat: 55.749321,
	//Lng: 37.608839,
	//},
	//ResultType: []string{
	//"street_address",
	//"route",
	//},
	//Language: "ru",
	//}
	//respOne, err := c.Geocode(context.Background(), rOne)
	//if err != nil {
	//log.Fatalf("fatal error: %s", err)
	//}

	//rTwo := &maps.GeocodingRequest{
	//LatLng: &maps.LatLng{
	//Lat: 55.754563,
	//Lng: 37.612369,
	//},
	//ResultType: []string{
	//"street_address",
	//"route",
	//},
	//Language: "ru",
	//}
	//respTwo, err := c.Geocode(context.Background(), rTwo)
	//if err != nil {
	//log.Fatalf("fatal error: %s", err)
	//}

	//pretty.Println(respOne)
	//pretty.Println(respTwo)
}
