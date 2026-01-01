package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bogem/id3v2/v2"
	"github.com/urfave/cli/v3"
)

func main() {
	if !(len(os.Args) > 1) {
		fmt.Println("No file name provided!")
		return
	}
	args := os.Args[1:]

	tag, err := id3v2.Open(args[0], id3v2.Options{Parse: true})
	if err != nil {
		fmt.Println("Error opening file, file may not exist or may not be an mp3 file.")
		return
	}
	defer tag.Close()

	var artistName string
	var albumName string
	var coverPath string
	var songTitle string

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "artist",
				Aliases: []string{"a", "ar"},
				Value:   "",
				Usage:   "Set artist name",
			},
			&cli.StringFlag{
				Name:    "title",
				Aliases: []string{"t"},
				Value:   "",
				Usage:   "Set song title",
			},
			&cli.StringFlag{
				Name:    "cover",
				Aliases: []string{"c", "image", "i"},
				Value:   "",
				Usage:   "Set cover image",
			},
			&cli.StringFlag{
				Name:    "album",
				Aliases: []string{"al"},
				Value:   "",
				Usage:   "Set album name",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() > 0 {
				artistName = cmd.String("artist")
				songTitle = cmd.String("title")
				coverPath = cmd.String("cover")
				albumName = cmd.String("album")
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

	if artistName != "" {
		tag.SetArtist(artistName)
	}

	if songTitle != "" {
		tag.SetTitle(songTitle)
	}

	if albumName != "" {
		tag.SetAlbum(albumName)
	}

	if coverPath != "" {
		info, err := os.Stat(coverPath)
		if err != nil || info.IsDir() {
			fmt.Println("Dropped item is not a valid file:", coverPath)
			return
		}

		data, err := os.ReadFile(coverPath)
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return
		}

		pic := id3v2.PictureFrame{
			Encoding:    id3v2.EncodingUTF8,
			MimeType:    "image/png",
			PictureType: id3v2.PTFrontCover,
			Description: "Cover",
			Picture:     data,
		}
		tag.AddAttachedPicture(pic)
	}

	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag: ", err)
	}

}
