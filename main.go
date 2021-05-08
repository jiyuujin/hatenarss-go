package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mmcdole/gofeed"
)

// GET /
func topPage(w http.ResponseWriter, r *http.Request) {
	fp := gofeed.NewParser()

	feed, _ := fp.ParseURL("https://b.hatena.ne.jp/entrylist/it.rss")
	items := feed.Items

	for _, item := range items {
		fmt.Fprintf(w, item.Title + "\n" + item.Link + "\n\n")
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println()
	}
}

func handleRequests() {
	http.HandleFunc("/", topPage)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func main() {
	handleRequests()
}
