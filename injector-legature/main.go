package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/net/context"
	"googlemaps.github.io/maps"

	"github.com/pik4ez/apisal/apisal"
)

// GeoCoder contains client to Google Maps API.
type GeoCoder struct {
	client *maps.Client
}

const apiKey = "AIzaSyCoTE7DPJAOPjauBDHeulN49M5aLc4QijY"

func main() {
	geocoder, err := NewGeoCoder(apiKey)
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
			if err != nil {
				continue
			}
			streetName := getStreetName(ungeo)
			if err != nil {
				continue
			}
			legatureObject, err := getLegature(streetName)
			if err == nil {
				objectsWriter.WriteObject(legatureObject)
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

func getLegature(objName string) (apisal.Object, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	legatures := map[int][]string{
		0: []string{
			"Начало пути: %s",
			"Точка старта: %s",
			"%s — то, что доктор прописал, чтобы начать маршрут",
		},
		1: []string{
			"Начинаем путь с %s",
			"Всё началось с %s",
			"Маршрут берёт своё начало от %s",
		},
		2: []string{
			"Сначала движемся по %s",
			"Старт маршрута -- вплотную к %s",
		},
		3: []string{
			"Находим %s, чтобы оттуда начать маршрут",
			"Берём %s как точку старта",
		},
		4: []string{
			"Стартуем рядом с %s",
			"Пробираемся %s к началу маршрута",
		},
		5: []string{
			"А на старте только и разговоров, что о %s",
		},
	}
	curCase := 0
	objNameCases, err := getObjNameCases(objName)
	if err != nil {
		return apisal.Object{}, err
	}
	legature := legatures[curCase][rand.Intn(len(legatures[curCase]))]
	legatureObject := apisal.Object{
		Type:        apisal.ObjectTypeLegature,
		Description: fmt.Sprintf(legature, objNameCases[curCase]),
	}
	return legatureObject, nil
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

func getObjNameCases(objName string) (map[int]string, error) {
	chars := []rune(objName)
	if strings.HasPrefix(objName, "улица") {
		body := string(chars[6:])
		return map[int]string{
			0: "улица " + body,
			1: "улицы " + body,
			2: "улице " + body,
			3: "улицу " + body,
			4: "улицей " + body,
			5: "улице " + body,
		}, nil
	}
	if strings.HasPrefix(objName, "переулок") {
		body := string(chars[9:])
		return map[int]string{
			0: "переулок " + body,
			1: "переулка " + body,
			2: "переулку " + body,
			3: "переулок " + body,
			4: "переулком " + body,
			5: "переулке " + body,
		}, nil
	}
	return nil, errors.New("no cases found")
}
