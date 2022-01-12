package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/msantocardoso/video-data-extractor/internal/core/domain"
	"github.com/msantocardoso/video-data-extractor/internal/core/ports"
)

type Video struct {
	repository ports.VideoRepository
}

func New(repository ports.VideoRepository) *Video {
	return &Video{
		repository: repository,
	}
}

func (v Video) Load(filename string) (*domain.Video, error) {

	video, err := v.repository.Get(filename)
	if err != nil {
		fmt.Printf("Load video usecase failed %+v", err)
		return nil, err
	}

	return video, nil
}

func (v Video) LoadAllFrom(path string) ([]*domain.Video, error) {
	var videos []*domain.Video
	extensions := v.loadExtensionMap()
	err := filepath.Walk(path, func(filename string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			valid := v.isValidExtension(filename, extensions)
			if valid {
				video, _ := v.repository.Get(filename)
				videos = append(videos, video)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Program failed on load root path ['" + path + "']!\n")
		return nil, err
	}

	fmt.Printf("Found [ %d ] files on [ '%s' ]\n", len(videos), path)

	return videos, nil
}

func (v Video) loadExtensionMap() map[string]bool {
	return map[string]bool{
		".mkv": true,
		".mp4": true,
		".avi": true,
	}
}

func (v Video) isValidExtension(filename string, extensions map[string]bool) bool {
	extension := filepath.Ext(filename)
	return extensions[extension]
}
