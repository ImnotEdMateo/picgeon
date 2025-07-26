package utils

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Media struct {
	Name     	string
	URL      	string
	ThumbURL	string
	IsVideo		bool
}

func ParseLinks(body io.Reader, baseURL string) ([]Media, error) {
	var media []Media
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
						isVideo = strings.HasSuffix(lower, ".mp4") || strings.HasSuffix(lower, ".webm")
						isImage = strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".jpeg") || strings.HasSuffix(lower, ".png")
					}
				}

				if isImage || isVideo {
					mediaItem := Media{
						Name:    href,
						URL:     baseURL + href,
						IsVideo: isVideo,
					}

					thumbName := strings.ReplaceAll(href, "/", "_")
					thumbPath, err := GetOrCreateThumbnail(mediaItem.URL, thumbName, mediaItem.IsVideo)
					if err == nil {
						mediaItem.ThumbURL = "/" + thumbPath
					} else {
						mediaItem.ThumbURL = mediaItem.URL
					}

					media = append(media, mediaItem)
				}
			}
		}
	}
}
