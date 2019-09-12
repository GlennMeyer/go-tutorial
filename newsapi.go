package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "os"

  "github.com/joho/godotenv"
)

type Source struct {
  Id          string
  Name        string
  Description string	
  Url         string
  Category    string
  Language    string
  Country     string
}

type Response struct {
  Status  string
  Sources []Source
}

func fetchSources() {
  var response Response
  err := godotenv.Load()
  if err != nil {
    panic(err)
  }
  apiKey := os.Getenv("NEWSAPI_KEY")
  resp, err := http.Get("https://newsapi.org/v2/sources?apiKey=" + apiKey)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  json.Unmarshal([]byte(body), &response)
  db := pgConnect()
  defer db.Close()
  _, err = db.Exec(`CREATE TABLE IF NOT EXISTS sources (
                      id varchar PRIMARY KEY,
                      name varchar,
                      description varchar,
                      url varchar,
                      category varchar,
                      language varchar,
                      country varchar
                    );`)
  if err != nil {
    panic(err)
  }
  if response.Status == "ok" {
    for _, source := range response.Sources {
      _, err := db.Model(&source).OnConflict("(id) DO UPDATE").Insert()
      if err != nil {
        panic(err)
      }
    }
  }
}
