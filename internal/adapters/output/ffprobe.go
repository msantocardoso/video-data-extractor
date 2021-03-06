package adapters

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type ProbeFormat struct {
	Filename         string            `json:"filename"`
	NBStreams        int               `json:"nb_streams"`
	NBPrograms       int               `json:"nb_programs"`
	FormatName       string            `json:"format_name"`
	FormatLongName   string            `json:"format_long_name"`
	StartTimeSeconds float64           `json:"start_time,string"`
	DurationSeconds  float64           `json:"duration,string"`
	Size             uint64            `json:"size,string"`
	BitRate          uint64            `json:"bit_rate,string"`
	ProbeScore       float64           `json:"probe_score"`
	Tags             map[string]string `json:"tags"`
}

func (f ProbeFormat) StartTime() time.Duration {
	return time.Duration(f.StartTimeSeconds * float64(time.Second))
}

func (f ProbeFormat) SizeInMB() uint64 {
	return f.Size / 1024 / 1024
}

func (f ProbeFormat) Duration() time.Duration {
	return time.Duration(f.DurationSeconds * float64(time.Second))
}

func (f ProbeFormat) DurationAsString() string {
	duration := f.Duration()
	return fmt.Sprintf("%d:%d:%.0f", int(duration.Hours()), int(duration.Minutes()), (duration % time.Minute).Seconds())
}

type ProbeData struct {
	Format *ProbeFormat `json:"format,omitempty"`
}

type ProbeAdapter struct{}

func (p ProbeAdapter) Get(filename string) (*ProbeData, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-show_format", "-print_format", "json=compact=1", filename)
	cmd.Stderr = os.Stderr

	r, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	var v ProbeData
	err = json.NewDecoder(r).Decode(&v)
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return &v, nil
}
