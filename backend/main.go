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
	URL   string
}

func main() {

	//Handling redirections
	http.HandleFunc("/current-season/", currentSeason)
	http.HandleFunc("/current-season/search/", currentSeasonSearch)

	log.Println("Listening..")
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func currentSeason(w http.ResponseWriter, r *http.Request) {
	shows := getCurrentSeason()

	fmt.Fprintf(w, toJSON(shows))
}

func currentSeasonSearch(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()[""]

	if !ok {

		fmt.Fprintf(w, "Search is empty!")

	} else {

		key := strings.ToLower(key[0])

		shows := getCurrentSeason()
		showsAfterQuery := []show{}
		for _, show := range shows {
			title := strings.ToLower(show.Title)
			if strings.Contains(title, key) {
				showsAfterQuery = append(showsAfterQuery, show)
			}
		}

		fmt.Fprintf(w, toJSON(showsAfterQuery))

	}
}

func getCurrentSeason() []show {
	shows := []show{}

	c := colly.NewCollector(
		colly.AllowedDomains("horriblesubs.info"),
	)

	//For every ind-show html element I parse it's title and href
	c.OnHTML(".ind-show", func(e *colly.HTMLElement) {
		temp := show{}
		temp.Title = e.ChildText("a[title]")
		temp.URL = e.ChildAttr("a[href]", "href")
		shows = append(shows, temp)
	})

	c.Visit("https://horriblesubs.info/current-season/")

	return shows
}

func toJSON(s []show) string {
	jsonBytes, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
