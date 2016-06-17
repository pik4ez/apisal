package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pik4ez/apisal/apisal"
)

func main() {
	if os.Stdin == nil {
		log.Fatal("sdfsdf")
	}
	points, err := AllPoints(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for _, point := range points {
		fmt.Println(point.Lat, point.Lon, point.Time)
	}
}

func AllPoints(r io.Reader) ([]apisal.Point, error) {
	var points []apisal.Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		point := apisal.Point{}
		err := json.Unmarshal([]byte(scanner.Text()), &point)
		if err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	return points, nil
}

//func Objects(point Point) ([]Object, error) {
//
//
//}
