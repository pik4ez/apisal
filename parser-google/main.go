package main

import (
	"os"
	"bufio"
	"fmt"
	"time"
	"encoding/json"
	"log"
)

type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Time time.Time `json:"time"`
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		point := Point{}
		err := json.Unmarshal([]byte(scanner.Text()), &point)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(point.Lat, point.Lon, point.Time)
	}
}
