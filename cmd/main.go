package main

import (
	adapters "github.com/msantocardoso/video-data-extractor/adapters/output"
	"github.com/msantocardoso/video-data-extractor/adapters/output/csv"
	"github.com/msantocardoso/video-data-extractor/adapters/repository"
	"github.com/msantocardoso/video-data-extractor/internal/core/usecase"
)

func main() {
	filename := "C:/Users/msant/Videos"
	videoRepository := repository.NewVideoRepository(&adapters.ProbeAdapter{})
	videoUsecase := usecase.New(videoRepository)

	videos, err := videoUsecase.LoadAllFrom(filename)
	if err != nil {
		panic("Program failed on load filename ['" + filename + "']!")
	}

	csvWriter := csv.New()

	csvWriter.Generate(filename, videos)
}
