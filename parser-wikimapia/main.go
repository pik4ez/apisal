package main

import (
	"os"
	"bufio"
	"fmt"
	"time"
	"encoding/json"
	"log"
	"io"
)

type Point struct {
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Time time.Time `json:"time"`
}

type Object struct {

}

type WikiObject struct {
	Title string `json:"title"`
}

func main() {
	if s, err := os.Stdin.Stat(); err != nil || s.Size() == 0 {
		log.Fatal("stdin is empty!")
	}

	points, err := AllPoints(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	for _, point := range points {
		fmt.Println(point.Lat, point.Lon, point.Time)
	}
}

func AllPoints(r io.Reader) ([]Point, error) {
	var points []Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		point := Point{}
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
