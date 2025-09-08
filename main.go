package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func fetchtitle(url string, ch chan<- string) {
	resp, err := http.Get(url) //sent HTTP get request to the url
	if err != nil {
		ch <- fmt.Sprintf("Error Fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error Parsing %s: %v", url, err)
		return
	}

	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc) //root node returned by html.Parse()

	if title == "" {
		title = "No Title found"
	}
	ch <- fmt.Sprintf("%s->%s", url, title)

}

func scrapeHandler(w http.ResponseWriter, r *http.Request) {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://www.youtube.com/",
		"https://www.geeksforgeeks.org",
	}

	ch := make(chan string)

	for _, url := range urls {
		go fetchtitle(url, ch)
	}

	for range urls {
		fmt.Fprintln(w, <-ch) // Print results to HTTP response
	}
}

func main() {
	http.HandleFunc("/scrape", scrapeHandler)

	fmt.Println("Server is running at http://localhost:8080/scrape")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
