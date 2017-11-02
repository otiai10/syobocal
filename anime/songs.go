package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/otiai10/anime"
	"github.com/urfave/cli"
)

// SongSearch ...
var SongSearch = cli.Command{
	Name:    "songs",
	Aliases: []string{"s"},
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

		songs := []anime.Song{}
		for _, anime := range animes {
			songs = append(songs, anime.Songs...)
		}

		if ctx.Bool("dump") {
			e := json.NewEncoder(os.Stdout)
			e.SetIndent("", "\t")
			return e.Encode(songs)
		}

		for _, song := range songs {
			fmt.Printf("%s / %s\n", song.Title, song.Anime)
		}

		return nil
	},
}
