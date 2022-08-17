package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/otiai10/syobocal/api"
	"github.com/otiai10/syobocal/factory"
)

func main() {
	if err := crawl(); err != nil {
		log.Fatalln(err)
	}
}

func crawl() error {

	// Define start/end time
	tokyo, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return fmt.Errorf("failed to load location: %v", err)
	}
	log.Println("Now:\t", time.Now())
	_24HoursAgo := time.Now().In(tokyo).Add(-24 * time.Hour)
	log.Println("24H Ago:\t", _24HoursAgo)
	start := time.Date(
		_24HoursAgo.Year(),
		_24HoursAgo.Month(),
		_24HoursAgo.Day(),
		_24HoursAgo.Hour(),
		0, 0, 0, tokyo,
	)
	end := start.Add(24 * time.Hour)
	log.Println("Start:\t", start)
	log.Println("End:\t", end)

	// Build query
	builder := api.TitleLookup().LastUpdate(start, end)
	query := builder.Build().Encode()
	log.Println("Query:\t", query)

	// HTTP request
	res, err := http.Get(api.BaseURL + "?" + query)
	log.Println("Actual:\t", res.Request.URL.String())
	if err != nil {
		return fmt.Errorf("failed to send GET request: %v", err)
	}
	if res.StatusCode >= 400 {
		return fmt.Errorf("http status seems ng: %v", res.Status)
	}
	defer res.Body.Close()

	// Parse XML
	lookup := api.TitleLookupResponse{}
	if err := xml.NewDecoder(res.Body).Decode(&lookup); err != nil {
		return fmt.Errorf("failed to decode XML: %v", err)
	}
	log.Println("Response:\t", lookup.Result.Code, lookup.Result.Message)
	if lookup.Result.Code != 200 {
		return fmt.Errorf("got syoboal non-200 response: %v %v", lookup.Result.Code, lookup.Result.Message)
	}

	// Convert to out model
	animes, err := factory.ToAnimeListFromTitleLookup(lookup)
	if err != nil {
		return fmt.Errorf("failed to convert to Anime struct: %v", err)
	}
	log.Println("Animes:\t", len(animes))

	// Prepare dir for JSON output
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current wd: %v", err)
	}
	dir := filepath.Join(pwd, "db", "historical")
	if err = os.MkdirAll(dir, os.ModeSticky|os.ModePerm); err != nil {
		return fmt.Errorf("failed to create dir for json output: %v", err)
	}

	// Write JSON file
	f, err := os.Create(filepath.Join(dir, start.Format("20060102")+".json"))
	if err != nil {
		return fmt.Errorf("failed to open file to write json output: %v", err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(animes); err != nil {
		return fmt.Errorf("failed to encode struct to json file: %v", err)
	}
	log.Println("File:\t", f.Name())

	return nil
}
