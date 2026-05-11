package scanning

import (
	"fmt"
	"os"
	"picgeon/utils"
)

type LocalScanner struct {
	Dir     string
	BaseURL string
}

func (s *LocalScanner) Scan() ([]utils.Media, error) {
	entries, err := os.ReadDir(s.Dir)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %s: %w", s.Dir, err)
	}

	var media []utils.Media
	for _, entry := range entries {
		if entry.IsDir() || !isMedia(entry.Name()) {
			continue
		}

		name := entry.Name()
		media = append(media, utils.Media{
			Name:    name,
			OrigURL: s.BaseURL + "/" + name,
			IsVideo: isVideo(name),
		})
	}

	return media, nil
}
