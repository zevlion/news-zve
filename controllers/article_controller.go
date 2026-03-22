package controllers

import (
	"net/http"
	"news-zve/services"

	"github.com/gin-gonic/gin"
)

func GetArticleDetail(c *gin.Context) {
	articleURL := c.Query("url")
	if articleURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL parameter is required"})
		return
	}

	detail, err := services.GetArticleDetail(articleURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}
