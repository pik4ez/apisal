package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gotokatsuya/gosquare/dispatcher"
	"github.com/gotokatsuya/gosquare/model"
	"github.com/gotokatsuya/gosquare/service/venues"
	"github.com/pik4ez/apisal/apisal"
)

func main() {
	if s, err := os.Stdin.Stat(); err != nil || s.Size() == 0 {
		log.Fatal("stdin is empty!")
	}

	points, err := ReadPoints(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	var venues, err = VenuesExplore()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range venues {
		fmt.Println(v.Name)
	}
}

// ReadPoints returns a list of points from stdin.
func ReadPoints(r io.Reader) ([]apisal.Point, error) {
	var points []apisal.Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		point := apisal.Point{}
		err := json.Unmarshal([]byte(scanner.Text()), &point)
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	return points, nil
}

// VenuesExplore returns venues near the specified location.
func VenuesExplore() ([]model.Venue, error) {
	client := dispatcher.NewClient()
	req := venues.NewExploreRequest()
	req.LatLng = "40.7,-74"
	res, err := venues.Explore(client, req)
	if err != nil {
		return nil, err
	}
	return res.GetVenues(), nil
}
