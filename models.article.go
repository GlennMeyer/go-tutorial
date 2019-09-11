package main

import (
  "errors"
)

type article struct {
  Id      int
  Title   string `sql:",notnull"`
  Content string `sql:",notnull"`
}

var articleList = []article{
  article{Id: 1, Title: "Article 1", Content: "Article 1 body"},
  article{Id: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []article {
  var articles []article
  db := pgConnect()
  defer db.Close()
  err := db.Model(&articles).Select()
  if err != nil {
    panic(err)
  }
  return articles
}

func getArticleByID(id int) (*article, error) {
  article := &article{Id: id}
  db := pgConnect()
  defer db.Close()
  err := db.Select(article)
  if err != nil {
    return nil, errors.New("Article not found")
  }
  return article, nil
}