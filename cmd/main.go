package main

import (
	adapters "github.com/msantocardoso/video-data-extractor/internal/adapters/output"
	"github.com/msantocardoso/video-data-extractor/internal/adapters/output/csv"
	"github.com/msantocardoso/video-data-extractor/internal/adapters/repository"
	"github.com/msantocardoso/video-data-extractor/internal/core/usecase"
)

func main() {
	filename := "C:/Users/msant/Videos/investalk-private-classes"
	videoRepository := repository.NewVideoRepository(&adapters.ProbeAdapter{})
	videoUsecase := usecase.New(videoRepository)

	videos, err := videoUsecase.LoadAllFrom(filename)
	if err != nil {
		panic("Program failed on load filename ['" + filename + "']!")
	}

	csvWriter := csv.New()

	csvWriter.Generate(filename, videos)
}
