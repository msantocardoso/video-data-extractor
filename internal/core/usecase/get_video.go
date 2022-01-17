package usecase

import (
	"fmt"

	adapters "github.com/msantocardoso/video-data-extractor/internal/adapters/output"
	"github.com/msantocardoso/video-data-extractor/internal/core/domain"
)

type Video struct {
	file_system_scan adapters.FileScan
}

func New(file_system_scan *adapters.FileScan) *Video {
	return &Video{
		file_system_scan: *file_system_scan,
	}
}

func (v Video) Load(filename string) (*domain.Video, error) {
	video, err := v.file_system_scan.Get(filename)
	if err != nil {
		fmt.Printf("Load video usecase failed %+v", err)
		return nil, err
	}

	return video, nil
}

func (v Video) LoadAllFrom(path string) ([]*domain.Video, error) {
	videos, err := v.file_system_scan.GetAllFrom(path)

	if err != nil {
		fmt.Printf("Program failed on load root path ['" + path + "']!\n")
		return nil, err
	}

	fmt.Printf("Found [ %d ] files on [ '%s' ]\n", len(videos), path)

	return videos, nil
}
