package controllers

import (
	"net/http"
	"news-zve/internal/cache"
	"news-zve/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var newsCache = cache.NewNewsCache()

func GetNews(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "25")
	limit, _ := strconv.Atoi(limitStr)

	cacheKey := "ap_news_" + limitStr

	if cachedArticles, found := newsCache.Get(cacheKey); found {
		c.Header("X-Cache", "HIT")
		c.JSON(http.StatusOK, cachedArticles)
		return
	}

	articles, err := services.FetchAPNews(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newsCache.Set(cacheKey, articles, 15*time.Minute)

	c.Header("X-Cache", "MISS")
	c.JSON(http.StatusOK, articles)
}
