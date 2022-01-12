package repository

import (
	"fmt"

	adapters "github.com/msantocardoso/video-data-extractor/internal/adapters/output"
	"github.com/msantocardoso/video-data-extractor/internal/core/domain"
	"github.com/msantocardoso/video-data-extractor/internal/core/ports"
)

type VideoRepository struct {
	probe *adapters.ProbeAdapter
}

func NewVideoRepository(probe *adapters.ProbeAdapter) ports.VideoRepository {
	return &VideoRepository{
		probe: probe,
	}
}

func (vr VideoRepository) Get(filename string) (*domain.Video, error) {
	result, err := vr.probe.Get(filename)

	if err != nil {
		fmt.Printf("Failed file load %s", filename)
		return nil, err
	}

	return domain.NewVideo(
		result.Format.Filename,
		result.Format.Size,
		result.Format.DurationAsString(),
	), nil
}
