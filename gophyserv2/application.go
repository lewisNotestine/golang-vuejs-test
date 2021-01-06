package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	giphyclient "gophyserv2/giphyClient"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	pageBody := string(p.Body)
	fmt.Fprintf(w, pageBody)
}

func gifHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Path[len("/api/gif/"):]
	gifUrl := giphyclient.GetGifUrl(input)
	fmt.Fprintf(w, gifUrl)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/api/gif/", gifHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
