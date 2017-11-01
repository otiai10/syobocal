package main

import (
	"fmt"
	"time"

	"github.com/otiai10/anime"
	"github.com/otiai10/anime/syobocal"
)

func main() {
	client := syobocal.NewClient()
	from := time.Now().Add(-17 * time.Hour)
	client.LastUpdated(&from, nil)
	res, err := client.Lookup()
	if err != nil {
		fmt.Println(err, client.Build())
		return
	}
	if len(res.TitleItems.Items) == 0 {
		return
	}
	info, err := anime.ParseComment(res.TitleItems.Items[0].Comment)
	fmt.Println(err)
	fmt.Printf("%+v\n", info)
}
