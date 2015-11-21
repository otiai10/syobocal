package main

import (
	"log"
	"time"

	"github.com/otiai10/animapi/factory"
	"github.com/otiai10/animapi/syobocal"
)

func main() {
	tlr, err := syobocal.NewClient().TitleLookup(time.Now().Add(-2*time.Minute), time.Now()).Do()
	// tlr, err := syobocal.NewClient().Do()
	if err != nil {
		panic(err)
	}
	if tlr.Result.Code != 200 {
		panic(tlr.Result.Message)
	}

	animes, err := factory.ConvertTitleLookupResponseToAnime(*tlr)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v\n", animes)
}
