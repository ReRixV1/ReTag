package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bogem/id3v2/v2"
	"github.com/urfave/cli/v3"
	"net.rerix/retag/internal/models"
	"net.rerix/retag/internal/services"
)

func main() {

	var fileInfo *models.FileInfo

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
			if !(cmd.Args().Len() > 0) {
				fmt.Println("No file name provided!")
				return nil
			}

			tag, err := id3v2.Open(cmd.Args().First(), id3v2.Options{Parse: true})
			if err != nil {
				fmt.Println("Error opening file, file may not exist or may not be an mp3 file.")
				return nil
			}
			defer tag.Close()

			if cmd.NArg() > 0 {
				fileInfo = &models.FileInfo{
					ArtistName: cmd.String("artist"),
					SongTitle:  cmd.String("title"),
					CoverPath:  cmd.String("cover"),
					AlbumName:  cmd.String("album"),
				}
			}

			if err := services.HandleSave(*fileInfo, tag); err != nil {
				fmt.Println("Error while saving file!")
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}

}
