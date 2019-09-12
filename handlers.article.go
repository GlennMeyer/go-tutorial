package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  articles := getAllArticles()
  c.HTML(http.StatusOK, "index.html", gin.H{ "title":   "Home Page", "payload": articles })
}

func getArticle(c *gin.Context) {
  if articleId, err := strconv.Atoi(c.Param("article_id")); err == nil {
    if article, err := getArticleByID(articleId); err == nil {
      c.HTML(http.StatusOK, "article.html", gin.H{ "title":   article.Title, "payload": article })
    } else {
      c.AbortWithError(http.StatusNotFound, err)
    }
  } else {
    c.AbortWithStatus(http.StatusNotFound)
  }
}