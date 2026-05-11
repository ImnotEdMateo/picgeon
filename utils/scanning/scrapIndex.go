package scanning

import (
	"io"
	"strings"

	"picgeon/utils"
	"golang.org/x/net/html"
)

func ParseLinks(body io.Reader, baseURL string) ([]utils.Media, error) {
	var media []utils.Media
	tokenizer := html.NewTokenizer(body)

	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return media, nil
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == "a" {
				var href string
				var isImage, isVideo bool

				for _, attr := range t.Attr {
					if attr.Key == "href" {
						href = attr.Val
						if href == "../" {
							continue
						}
						lower := strings.ToLower(href)
						isVideo = strings.HasSuffix(lower, ".mp4") || strings.HasSuffix(lower, ".webm") || strings.HasSuffix(lower, ".gif")
						isImage = strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".jpeg") || strings.HasSuffix(lower, ".png")
					}
				}

				if isImage || isVideo {
					media = append(media, utils.Media{
						Name:    href,
						URL:    baseURL + href,
						IsVideo: isVideo,
					})
				}
			}
		}
	}
}
