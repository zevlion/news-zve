package cache

import (
	"news-zve/models"
	"sync"
	"time"
)

type CacheItem struct {
	Articles  []models.Article
	ExpiresAt time.Time
}

type NewsCache struct {
	mu    sync.RWMutex
	store map[string]CacheItem
}

func NewNewsCache() *NewsCache {
	return &NewsCache{
		store: make(map[string]CacheItem),
	}
}

func (c *NewsCache) Set(key string, articles []models.Article, duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheItem{
		Articles:  articles,
		ExpiresAt: time.Now().Add(duration),
	}
}

func (c *NewsCache) Get(key string) ([]models.Article, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.store[key]
	if !found || time.Now().After(item.ExpiresAt) {
		return nil, false
	}
	return item.Articles, true
}
