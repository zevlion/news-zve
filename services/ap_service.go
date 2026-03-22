package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"news-zve/models"
	"time"
)

const (
	AP_ROOT_UUID = "00000000-0000-4000-8000-3caf4df03307"
	BASE_URL     = "https://livecomments.viafoura.co/v4/livecomments"
)

func FetchAPNews(limit int) ([]models.Article, error) {
	apiURL := fmt.Sprintf("%s/%s/trending?limit=%d&content_container_window_days=7&content_window_hours=1&sorted_by=total_visible_contents", BASE_URL, AP_ROOT_UUID, limit)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	req.Header.Set("Referer", "https://apnews.com/")
	req.Header.Set("Origin", "https://apnews.com")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api status error: %d", resp.StatusCode)
	}

	var data models.ViafouraResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var articles []models.Article
	for _, item := range data.Trending {
		articles = append(articles, models.Article{
			Title:     item.OriginTitle,
			Summary:   item.OriginSummary,
			URL:       item.OriginURL,
			Thumbnail: item.OriginImageURL,
			Source:    "AP News",
		})
	}

	return articles, nil
}
