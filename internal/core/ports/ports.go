package ports

import "github.com/msantocardoso/video-data-extractor/internal/core/domain"

type VideoRepository interface {
	Get(filename string) (*domain.Video, error)
}
