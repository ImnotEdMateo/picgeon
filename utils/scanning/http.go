package scanning

import (
	"fmt"
	"net/http"
	"picgeon/utils"
)

type HTTPScanner struct {
	BaseURL string
}

func (s *HTTPScanner) Scan() ([]utils.Media, error) {
	resp, err := http.Get(s.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching index: %w", err)
	}
	defer resp.Body.Close()

	return parseLinks(resp.Body, s.BaseURL)
}
