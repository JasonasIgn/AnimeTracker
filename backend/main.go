// test
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

type Show struct {
	Title       string
	URL         string
	Cover       string
	Description string
}

const mainSite = "https://horriblesubs.info"

func main() {

	//Handling END-POINTS
	http.HandleFunc("/current-season/", currentSeason)
	http.HandleFunc("/current-season/search/", currentSeasonSearch)

	log.Println("Listening..")
	log.Fatal(http.ListenAndServe(":1337", nil))
}

func currentSeason(w http.ResponseWriter, r *http.Request) {
	shows := currentSeasonDetails(getCurrentSeasonUrls())

	fmt.Fprintf(w, toJSON(shows))
}

func currentSeasonSearch(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()[""]

	if !ok {

		fmt.Fprintf(w, "Search is empty!")

	} else {

		key := strings.ToLower(key[0])

		shows := currentSeasonDetails(getCurrentSeasonUrls())
		showsAfterQuery := []Show{}
		for _, show := range shows {
			title := strings.ToLower(show.Title)
			if strings.Contains(title, key) {
				showsAfterQuery = append(showsAfterQuery, show)
			}
		}

		fmt.Fprintf(w, toJSON(showsAfterQuery))

	}
}

func getCurrentSeasonUrls() []string {
	showURLs := []string{}

	c := colly.NewCollector(
		colly.AllowedDomains("horriblesubs.info"),
		colly.AllowURLRevisit(),
		//colly.Async(true),
	)

	//For every ind-show html element I parse it's title and href
	c.OnHTML(".ind-show", func(e *colly.HTMLElement) {
		showURL := e.ChildAttr("a[href]", "href")
		showURLs = append(showURLs, showURL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping.")

	})

	c.Visit("https://horriblesubs.info/current-season/")

	return showURLs
}

func currentSeasonDetails(urls []string) []Show {
	var listOfShows = []Show{}

	detailCollector := colly.NewCollector()

	q, _ := queue.New(
		20, //Consumer threads (STILL NEED TO WORK THIS OUT)
		&queue.InMemoryQueueStorage{MaxSize: 10000}, //Size of queue
	)

	detailCollector.OnHTML(".site-content", func(e *colly.HTMLElement) {
		temp := Show{}
		temp.Title = e.ChildText(".entry-title")
		temp.URL = e.Request.URL.String()
		temp.Cover = e.ChildAttr("div.series-image img[src^='']", "src")
		temp.Description = e.ChildText(".series-desc")[16:]
		listOfShows = append(listOfShows, temp)
	})

	for _, showURL := range urls {
		q.AddURL(mainSite + showURL)
	}

	q.Run(detailCollector)

	return listOfShows
}

func toJSON(s []Show) string {
	jsonBytes, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
