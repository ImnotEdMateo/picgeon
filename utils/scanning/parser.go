package scanning

import (
	"io"

	"golang.org/x/net/html"
	"picgeon/utils"
)

func parseLinks(body io.Reader, baseURL string) ([]utils.Media, error) {
	var media []utils.Media
	tokenizer := html.NewTokenizer(body)

	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return media, nil
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data != "a" {
				continue
			}

			var href string
			for _, attr := range t.Attr {
				if attr.Key == "href" {
					href = attr.Val
					break
				}
			}

			if href == "" || href == "../" || !isMedia(href) {
				continue
			}

			media = append(media, utils.Media{
				Name:    href,
				OrigURL: baseURL + href,
				IsVideo: isVideo(href),
			})
		}
	}
}
