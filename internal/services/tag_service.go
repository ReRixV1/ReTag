package services

import (
	"fmt"
	"os"

	"github.com/bogem/id3v2/v2"
	"net.rerix/retag/internal/models"
)

func setIfAvailible(fileInfo models.FileInfo, tag *id3v2.Tag) {
	if fileInfo.ArtistName != "" {
		tag.SetArtist(fileInfo.ArtistName)
	}

	if fileInfo.SongTitle != "" {
		tag.SetTitle(fileInfo.SongTitle)
	}

	if fileInfo.AlbumName != "" {
		tag.SetAlbum(fileInfo.AlbumName)
	}

}

func HandleSave(fileInfo models.FileInfo, tag *id3v2.Tag) error {
	setIfAvailible(fileInfo, tag)
	if fileInfo.CoverPath != "" {
		info, err := os.Stat(fileInfo.CoverPath)
		if err != nil || info.IsDir() {
			fmt.Println("Dropped item is not a valid file:", fileInfo.CoverPath)
			return nil
		}

		data, err := os.ReadFile(fileInfo.CoverPath)
		if err != nil {
			fmt.Println("Failed to read file:", err)
			return nil
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

	if err := tag.Save(); err != nil {
		return err
	}
	return nil
}
