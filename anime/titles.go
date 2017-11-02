package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/otiai10/anime"
	"github.com/urfave/cli"
)

// TitleSearch ...
var TitleSearch = cli.Command{
	Name:    "titles",
	Aliases: []string{"t"},
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name: "dump,D",
		},
	},
	Action: func(ctx *cli.Context) error {
		client := anime.NewClient()
		animes, err := client.Lookup()
		if err != nil {
			return err
		}

		if ctx.Bool("dump") {
			e := json.NewEncoder(os.Stdout)
			e.SetIndent("", "\t")
			return e.Encode(animes)
		}

		for _, anime := range animes {
			fmt.Println(anime.LastUpdated.Format("2006/01/02 15:04"), anime.Title)
		}

		return nil
	},
}
