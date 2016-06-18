package mapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const MAPIA_API_URL = "http://api.wikimapia.org/"

type Mapia struct {
	ApiKey string
}

type MapiaParams map[string]interface{}

// http://json2struct.mervine.net/
type MapiaPlaces struct {
	Count    int                 `json:"count"`
	Found    int                 `json:"found"`
	Language string              `json:"language"`
	Places   []MapiaPlaceSummary `json:"places"`
}

type MapiaPlaceSummary struct {
	Distance   int    `json:"distance"`
	ID         int    `json:"id"`
	Language   string `json:"language"`
	LanguageID int    `json:"language_id"`
	Location   struct {
		East  float64 `json:"east"`
		Lat   float64 `json:"lat"`
		Lon   float64 `json:"lon"`
		North float64 `json:"north"`
		South float64 `json:"south"`
		West  float64 `json:"west"`
	} `json:"location"`
	Title   string `json:"title"`
	Urlhtml string `json:"urlhtml"`
}

type MapiaPlace struct {
	Comments     []interface{} `json:"comments"`
	Description  string        `json:"description"`
	ID           int           `json:"id"`
	IsBuilding   bool          `json:"is_building"`
	IsDeleted    bool          `json:"is_deleted"`
	IsRegion     bool          `json:"is_region"`
	LanguageID   int           `json:"language_id"`
	LanguageIso  string        `json:"language_iso"`
	LanguageName string        `json:"language_name"`
	ObjectType   int           `json:"object_type"`
	ParentID     string        `json:"parent_id"`
	Photos       []struct {
		One280URL          string `json:"1280_url"`
		Nine60URL          string `json:"960_url"`
		BigURL             string `json:"big_url"`
		FullURL            string `json:"full_url"`
		ID                 int    `json:"id"`
		LastUserID         int    `json:"last_user_id"`
		LastUserName       string `json:"last_user_name"`
		ObjectID           int    `json:"object_id"`
		Size               int    `json:"size"`
		Status             int    `json:"status"`
		ThumbnailRetinaURL string `json:"thumbnailRetina_url"`
		ThumbnailURL       string `json:"thumbnail_url"`
		Time               int    `json:"time"`
		TimeStr            string `json:"time_str"`
		UserID             int    `json:"user_id"`
		UserName           string `json:"user_name"`
	} `json:"photos"`
	Pl   float64 `json:"pl"`
	Tags []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"tags"`
	Title   string `json:"title"`
	Urlhtml string `json:"urlhtml"`
	X       string `json:"x"`
	Y       string `json:"y"`
}

func NewMapia(apiKey string) *Mapia {

	mapia := &Mapia{
		ApiKey: apiKey,
	}

	return mapia
}

func (s *Mapia) GetPlaceById(id int, language string) (*MapiaPlace, error) {

	raw, err := s.CallApi("place.getbyid", &MapiaParams{
		"id":          id,
		"data_blocks": "main,photos,comments",
		"language":    language,
	})

	if err != nil {
		return nil, err
	}

	var place MapiaPlace

	// fmt.Println(raw)

	if err := json.Unmarshal([]byte(raw), &place); err != nil {
		return nil, err
	}

	return &place, nil
}

func (s *Mapia) GetNearbyObjects(lat, lon float64, language string) (*MapiaPlaces, error) {

	raw, err := s.CallApi("place.getnearest", &MapiaParams{
		"lat":         lat,
		"lon":         lon,
		"data_blocks": "location",
		"language":    language,
		"count":       9,
	})

	if err != nil {
		return nil, err
	}

	var places MapiaPlaces

	if err := json.Unmarshal([]byte(raw), &places); err != nil {
		return nil, err
	}

	return &places, nil
}

func (s *Mapia) CallApi(method string, params *MapiaParams) (string, error) {

	url, err := url.Parse(fmt.Sprintf("%s", MAPIA_API_URL))

	if err != nil {
		panic(err)
	}

	query := url.Query()

	query.Set("key", s.ApiKey)
	query.Set("function", method)
	query.Set("format", "json")

	if params != nil {
		for k, v := range *params {
			query.Set(k, fmt.Sprintf("%v", v))
		}
	}

	url.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Timeout = 10 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Timeout")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
