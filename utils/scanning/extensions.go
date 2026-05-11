package scanning

import "strings"

func isImage(name string) bool {
	lower := strings.ToLower(name)
	return strings.HasSuffix(lower, ".jpg") ||
		strings.HasSuffix(lower, ".jpeg") ||
		strings.HasSuffix(lower, ".png")
}

func isVideo(name string) bool {
	lower := strings.ToLower(name)
	return strings.HasSuffix(lower, ".mp4") ||
		strings.HasSuffix(lower, ".webm") ||
		strings.HasSuffix(lower, ".gif")
}

func isMedia(name string) bool {
	return isImage(name) || isVideo(name)
}
