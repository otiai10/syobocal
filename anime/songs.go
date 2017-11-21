package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/otiai10/anime"
	"github.com/otiai10/jsonindent"
	"github.com/urfave/cli"
)

// SongSearch ...
var SongSearch = cli.Command{
	Name:    "songs",
	Aliases: []string{"s"},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format,f",
			Usage: "Output format, either [json, csv, tsv] or [pretty] as default",
		},
		cli.StringFlag{
			Name:  "since,s",
			Usage: "Start of time to search",
		},
		cli.StringFlag{
			Name:  "until,u",
			Usage: "End of time to search",
		},
		cli.BoolFlag{
			Name:  "verbose,v",
			Usage: "Show verbose log",
		},
		cli.StringFlag{
			Name:  "timeformat",
			Usage: "Time format for `since` and `until`",
			Value: "YYYY-MM-dd",
		},
	},
	Action: func(ctx *cli.Context) error {

		timeformat := strings.Replace(ctx.String("timeformat"), "YYYY", "2006", -1)
		timeformat = strings.Replace(timeformat, "MM", "01", -1)
		timeformat = strings.Replace(timeformat, "dd", "02", -1)

		term := []time.Time{}
		if s := ctx.String("since"); s != "" {
			t, err := time.Parse(timeformat, s)
			if err != nil {
				return err
			}
			term = append(term, t)
		}
		if u := ctx.String("until"); u != "" && len(term) != 0 {
			t, err := time.Parse(timeformat, u)
			if err != nil {
				return err
			}
			term = append(term, t)
		}

		client := anime.NewClient()
		client.Verbose = ctx.Bool("verbose")
		animes, err := client.Lookup(term...)
		if err != nil {
			return err
		}

		songs := []anime.Song{}
		for _, anime := range animes {
			songs = append(songs, anime.Songs...)
		}

		switch ctx.String("format") {
		case "json":
			return jsonindent.NewEncoder(os.Stdout).Encode(songs)
		case "csv":
			return outputAsRowOrientedDataFormat(songs, "%[1]s,%[2]s,%[3]s")
		case "tsv":
			return outputAsRowOrientedDataFormat(songs, "%[1]s\t%[2]s\t%[3]s")
		default:
			return outputAsRowOrientedDataFormat(songs, "%[1]s (%[2]s) / %[3]s")
		}

	},
}

func outputAsRowOrientedDataFormat(songs []anime.Song, format string) error {
	for _, song := range songs {
		fmt.Printf(format+"\n", song.Anime, song.Label, song.Title)
	}
	return nil
}
