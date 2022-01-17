package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/msantocardoso/video-data-extractor/internal/core/domain"
)

type Csv struct {
	filename string
}

func New(filename string) *Csv {
	return &Csv{filename: filename}
}

func (c Csv) Generate(path string, videos []*domain.Video) {
	file, err := os.Create(path + "/" + c.filename + ".csv")

	defer file.Close()
	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	w := csv.NewWriter(file)

	defer w.Flush()

	for _, video := range videos {
		refers := strings.Split(video.FileName(), "_")

		modNameWithoutPath := strings.Split(refers[0], "\\")
		modName := strings.ReplaceAll(modNameWithoutPath[len(modNameWithoutPath)-1], "-", " ")
		className := strings.ReplaceAll(refers[1], "-", " ")
		dateAndExtension := strings.Split(refers[2], ".")
		date := dateAndExtension[0]
		extension := dateAndExtension[1]

		row := []string{
			date,
			modName,
			className,
			extension,
			video.FileName(),
			video.Duration,
			strconv.FormatUint(video.Size(), 10),
		}

		if err := w.Write(row); err != nil {
			log.Fatalln("Error when writing record to file", err)
		}
	}

}
