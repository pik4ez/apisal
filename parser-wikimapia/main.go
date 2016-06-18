package main

import (
	"os"
	"bufio"
	"encoding/json"
	"log"
	"io"
	lib "github.com/pik4ez/apisal/apisal"
	//"fmt"
	mapia "./mapia"
	"fmt"
)

const API_KEY = "59F5F0FD-B38A4635-6BC3D0EF-307471CE-7246D42E-A54DC82A-BB2A27C6-9A0FD0BE"

func main() {
	if s, err := os.Stdin.Stat(); err != nil || s.Size() == 0 {
		log.Fatal("stdin is empty!")
	}

	points, err := AllPoints(os.Stdin)
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

	for _, object := range objects {
		str, err := json.Marshal(object)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(str))
	}
}

func AllPoints(r io.Reader) ([]lib.Point, error) {
	var points []lib.Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		point := lib.Point{}
		err := json.Unmarshal([]byte(scanner.Text()), &point)
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	return points, nil
}

func PointObjects(point lib.Point) ([]lib.Object, error) {
	var objects []lib.Object

	m := mapia.NewMapia(API_KEY)
	places, err := m.GetNearbyObjects(point.Lat, point.Lon, "ru")
	if err != nil {
		return nil, err
	}

	// fmt.Println(places.Count, places.Found, places.Language)

	for _, place := range places.Places {
		extened, err := m.GetPlaceById(place.ID, "ru")
		if err != nil {
			return nil, err
		}
		images := make([]lib.Image, 0)
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
