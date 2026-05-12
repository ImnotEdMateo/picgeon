package store

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"picgeon/utils"
	"picgeon/utils/scanning"
)

type GalleryStore struct {
	mu        sync.RWMutex
	items     []utils.Media
	lastFetch time.Time
	ttl       time.Duration
	scanner   scanning.Scanner
}

var Default *GalleryStore

func NewStore(scanner scanning.Scanner, ttl time.Duration) *GalleryStore {
	return &GalleryStore{
		scanner: scanner,
		ttl:     ttl,
	}
}

func (s *GalleryStore) Get() ([]utils.Media, error) {
	s.mu.RLock()
	if s.items != nil && time.Since(s.lastFetch) < s.ttl {
		items := s.items
		s.mu.RUnlock()
		return items, nil
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	// double-check: otro goroutine pudo haber refrescado mientras esperábamos
	if s.items != nil && time.Since(s.lastFetch) < s.ttl {
		return s.items, nil
	}

	items, err := s.scanner.Scan()
	if err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	s.items = enrich(items)
	s.lastFetch = time.Now()
	return s.items, nil
}

func (s *GalleryStore) Invalidate() {
	s.mu.Lock()
	s.lastFetch = time.Time{}
	s.mu.Unlock()
}

func enrich(items []utils.Media) []utils.Media {
	for i, item := range items {
		thumbName := strings.ReplaceAll(item.Name, "/", "_")
		thumbPath, err := utils.GetOrCreateThumbnail(item.OrigURL, thumbName, item.IsVideo)
		if err == nil {
			items[i].ThumbURL = "/" + thumbPath
		} else {
			items[i].ThumbURL = item.OrigURL
		}
	}
	return items
}
