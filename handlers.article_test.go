package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
  "fmt"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
  r := getRouter(true)

  r.GET("/", showIndexPage)

  req, _ := http.NewRequest("GET", "/", nil)

  testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
    statusOK := w.Code == http.StatusOK

    p, err := ioutil.ReadAll(w.Body)
    pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

    return statusOK && pageOK
  })
}

func TestArticleUnauthenticated(t *testing.T) {
  r := getRouter(true)

  r.GET("/article/view/:article_id", getArticle)

  for _, a := range articleList {
    req, _ := http.NewRequest("GET", fmt.Sprintf("/article/view/%d", a.ID), nil)

    testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
      statusOK := w.Code == http.StatusOK
  
      p, err := ioutil.ReadAll(w.Body)
      pageOK := err == nil && strings.Index(string(p), fmt.Sprintf("<h1>Article %d</h1>", a.ID)) > 0
  
      return statusOK && pageOK
    })
  }
}