package models

type Article struct {
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnail"`
	Source    string `json:"source"`
}

type ArticleContent struct {
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Date    string   `json:"date"`
	Content string `json:"content"`
}

type ViafouraResponse struct {
	Trending []struct {
		OriginTitle    string `json:"origin_title"`
		OriginSummary  string `json:"origin_summary"`
		OriginURL      string `json:"origin_url"`
		OriginImageURL string `json:"origin_image_url"`
	} `json:"trending"`
}
