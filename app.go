package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type TrendingArticles struct {
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      string `json:"source>name"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	Content     string `json:"content"`
}

type ArticleAggPage struct {
	Title    string
	Articles TrendingArticles
}

func cleanup() {
	r := recover()
	if r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	defer cleanup()
	var t TrendingArticles
	resp, _ := http.Get("https://newsapi.org/v2/top-headlines?country=us&apiKey=5b99e9a609584c1fb1f13793c084f410")
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	json.Unmarshal(bytes, &t)
	p := ArticleAggPage{Title: "Trending Articles on Run", Articles: t}
	ex, _ := template.ParseFiles("homePage.html")
	ex.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to golang!")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/Articles/", newsAggHandler)
	http.ListenAndServe(":3010", nil)
}
