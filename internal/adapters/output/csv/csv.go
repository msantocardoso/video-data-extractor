package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/msantocardoso/video-data-extractor/internal/core/domain"
)

type Csv struct{}

func New() *Csv {
	return &Csv{}
}

func (c Csv) Generate(path string, videos []*domain.Video) {
	now := strings.ReplaceAll(time.Now().Format(time.RFC3339), ":", "")
	file, err := os.Create(path + "/all-videos-" + now + ".csv")

	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(file)

	defer w.Flush()

	for _, video := range videos {
		row := []string{
			video.FileName(),
			video.Duration,
			strconv.FormatUint(video.Size(), 10),
		}

		if err := w.Write(row); err != nil {
			log.Fatalln("Error when writing record to file", err)
		}
	}

}
