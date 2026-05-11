package utils

type Media struct {
	Name			string		`json:"name"`
	OrigURL		string		`json:"url"`
	ThumbURL	string		`json:"thumb_url"`
	IsVideo		bool			`json:"is_video"`
}
