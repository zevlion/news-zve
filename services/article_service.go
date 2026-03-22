package services

import (
	"fmt"
	"net/http"
	"news-zve/models"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetArticleDetail(url string) (*models.ArticleContent, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("api error: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	article := &models.ArticleContent{}
	article.Title = strings.TrimSpace(doc.Find("h1").First().Text())

	tsStr, exists := doc.Find("bsp-timestamp").Attr("data-timestamp")
	if exists {
		ms, err := strconv.ParseInt(tsStr, 10, 64)
		if err == nil {
			t := time.Unix(0, ms*int64(time.Millisecond))
			article.Date = t.Format("January 02, 2006 at 3:04 PM")
		}
	}

	if article.Date == "" {
		article.Date = strings.TrimSpace(doc.Find("span[data-date]").First().Text())
	}

	doc.Find(".Page-authors a").Each(func(i int, s *goquery.Selection) {
		name := strings.TrimSpace(s.Text())
		if name != "" {
			article.Authors = append(article.Authors, name)
		}
	})

	var paragraphs []string
	doc.Find(".RichTextStoryBody p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if i == 0 && strings.Contains(text, " (AP) — ") {
			parts := strings.SplitN(text, " (AP) — ", 2)
			if len(parts) > 1 {
				text = parts[1]
			}
		}
		if text != "" {
			paragraphs = append(paragraphs, text)
		}
	})

	article.Content = strings.Join(paragraphs, "\n\n")

	return article, nil
}
