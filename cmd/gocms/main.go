package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("/home/disbeliefff/GolandProjects/MyPortfolioWeb/templates /index.tmpl")

	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Home",
		})
	})

	r.Static("/templates", "/home/disbeliefff/GolandProjects/MyPortfolioWeb/templates /index.tmpl")
	r.Static("/templates", "/home/disbeliefff/GolandProjects/MyPortfolioWeb/templates /index.css")
	r.Run()
}
