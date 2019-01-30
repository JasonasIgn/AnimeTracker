// test
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Show struct {
	Title       string
	URL         string
	Cover       string
	Description string
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

func getCurrentSeason() []Show {
	ch := make(chan Show)

	go func() {

		const mainSite = "https://horriblesubs.info"

		c := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0"),
			colly.AllowedDomains("horriblesubs.info"),
			colly.Async(true),
		)

		c.Limit(&colly.LimitRule{
			DomainGlob:  ".*horriblesubs.*",
			Parallelism: 1,
			Delay:       1 * time.Second,
		})

		detailCollector := c.Clone()

		//For every ind-show html element I parse it's title and href
		c.OnHTML(".ind-show", func(e *colly.HTMLElement) {
			showUrl := e.ChildAttr("a[href]", "href")
			detailCollector.Visit(mainSite + showUrl)

		})

		detailCollector.OnHTML(".site-content", func(e *colly.HTMLElement) {

			title := e.ChildText(".entry-title")
			url := e.Request.URL.String()
			cover := e.ChildAttr("div.series-image img[src^='']", "src")
			description := e.ChildText(".series-desc")[16:]
			ch <- Show{Title: title, URL: url, Cover: cover, Description: description}
			time.Sleep(time.Second * 1)
		})

		c.Visit("https://horriblesubs.info/current-season/")
		c.Wait()
		close(ch)

	}()

	shows := []Show{}

	for show := range ch {
		fmt.Println(show.Title)
		shows = append(shows, show)

	}

	return shows
}

func toJSON(s []Show) string {
	jsonBytes, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
