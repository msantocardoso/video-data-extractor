package adapters

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/msantocardoso/video-data-extractor/internal/core/domain"
)

type FileScan struct {
	probe              ProbeAdapter
	allowed_extensions map[string]bool
}

func NewFileScan(allowed_extensions map[string]bool, probe *ProbeAdapter) *FileScan {
	return &FileScan{
		probe:              *probe,
		allowed_extensions: allowed_extensions,
	}
}

func (fs FileScan) GetAllFrom(path string) ([]*domain.Video, error) {
	var videos []*domain.Video
	err := filepath.Walk(path, func(filename string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			valid := isValidExtension(filename, fs.allowed_extensions)
			if valid {
				video, _ := fs.Get(filename)
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

func (fs FileScan) Get(filename string) (*domain.Video, error) {
	result, err := fs.probe.Get(filename)

	if err != nil {
		fmt.Printf("Failed file load %s", filename)
		return nil, err
	}

	return domain.NewVideo(
		result.Format.Filename,
		result.Format.SizeInMB(),
		result.Format.DurationAsString(),
	), nil
}

func isValidExtension(filename string, extensions map[string]bool) bool {
	extension := filepath.Ext(filename)
	return extensions[extension]
}
