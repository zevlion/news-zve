package main

import (
	"github.com/gin-gonic/gin"
	"news-zve/controllers"
)

func main() {
	r := gin.Default()

	r.StaticFile("/favicon.ico", "./media/favicon.ico")
	r.Static("/media", "./media")
	r.StaticFile("/", "./html/index.html")
	r.Static("/html", "./html")

	r.GET("/api/v1/news", controllers.GetNews)
	r.GET("/api/v1/article", controllers.GetArticleDetail)

	r.Run(":8080")
}
