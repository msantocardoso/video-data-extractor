package main

import (
	"strings"
	"time"

	adapters "github.com/msantocardoso/video-data-extractor/internal/adapters/output"
	"github.com/msantocardoso/video-data-extractor/internal/adapters/output/csv"
	"github.com/msantocardoso/video-data-extractor/internal/core/usecase"
)

func main() {
	path := "C:/Users/msant/Videos/investalk-private-classes"

	allowExtensions := map[string]bool{
		".mkv": true,
		".mp4": true,
		".avi": true,
	}
	filescan := adapters.NewFileScan(allowExtensions, &adapters.ProbeAdapter{})
	videoUsecase := usecase.New(filescan)

	videos, err := videoUsecase.LoadAllFrom(path)
	if err != nil {
		panic("Program failed on load filename ['" + path + "']!")
	}

	now := strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "")
	nowUnformatted := strings.ReplaceAll(now, "-", "")
	filename := "all-videos-" + nowUnformatted

	csvWriter := csv.New(filename)

	csvWriter.Generate(path, videos)
}
