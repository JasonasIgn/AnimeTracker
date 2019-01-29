// test
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
)

type show struct {
	Title string
	Url   string
}

func main() {

	//Handling redirections
	http.HandleFunc("/current-season/", current_season)
	http.HandleFunc("/current-season/search/", current_season_search)

	log.Println("Listening..")
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func current_season(w http.ResponseWriter, r *http.Request) {
	shows := get_current_season()

	fmt.Fprintf(w, to_json(shows))
}

func current_season_search(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()[""]

	if !ok {

		fmt.Fprintf(w, "Search is empty!")

	} else {

		key := strings.ToLower(key[0])

		shows := get_current_season()
		shows_after_query := []show{}
		for _, show := range shows {
			title := strings.ToLower(show.Title)
			if strings.Contains(title, key) {
				shows_after_query = append(shows_after_query, show)
			}
		}

		fmt.Fprintf(w, to_json(shows_after_query))

	}
}

func get_current_season() []show {
	shows := []show{}

	c := colly.NewCollector(
		colly.AllowedDomains("horriblesubs.info"),
	)

	c.OnHTML(".ind-show", func(e *colly.HTMLElement) {
		temp := show{}
		temp.Title = e.ChildText("a[title]")
		temp.Url = e.ChildAttr("a[href]", "href")
		//		fmt.Printf("FOUND -> %s @ %s\n", temp.Title, temp.ShowURL)
		shows = append(shows, temp)
	})

	//	c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL.String())
	//	})

	c.Visit("https://horriblesubs.info/current-season/")
	//	for index, show := range shows {
	//		fmt.Printf("%d. %s @ %s\n", index+1, show.Title, show.URL)
	//	}

	return shows
}

func to_json(s []show) string {
	json_bytes, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	return string(json_bytes)
}
