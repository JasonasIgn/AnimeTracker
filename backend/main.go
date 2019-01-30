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
	showURLS := getCurrentSeasonUrls()
	shows := currentSeasonDetails(showURLS)
	shows2 := []Show{}

	for show := range shows {
		fmt.Println(show.Title)
		shows2 = append(shows2, show)

	}

	fmt.Fprintf(w, toJSON(shows2))
}

func currentSeasonSearch(w http.ResponseWriter, r *http.Request) {
	key, ok := r.URL.Query()[""]

	if !ok {

		fmt.Fprintf(w, "Search is empty!")

	} else {

		key := strings.ToLower(key[0])

		showURLS := getCurrentSeasonUrls()
		shows := currentSeasonDetails(showURLS)
		showsAfterQuery := []Show{}
		for show := range shows {
			title := strings.ToLower(show.Title)
			if strings.Contains(title, key) {
				showsAfterQuery = append(showsAfterQuery, show)
			}
		}

		fmt.Fprintf(w, toJSON(showsAfterQuery))

	}
}

func getCurrentSeasonUrls() chan string {
	ch := make(chan string)

	go func() {
		const mainSite = "https://horriblesubs.info"

		c := colly.NewCollector(
			colly.AllowedDomains("horriblesubs.info"),
			colly.AllowURLRevisit(),
			colly.Async(true),
		)

		//For every ind-show html element I parse it's title and href
		c.OnHTML(".ind-show", func(e *colly.HTMLElement) {
			showUrl := e.ChildAttr("a[href]", "href")
			ch <- mainSite + showUrl
		})

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})

		c.OnScraped(func(response *colly.Response) {
			fmt.Println("LOL")

		})

		c.Visit("https://horriblesubs.info/current-season/")

	}()

	return ch
}

func currentSeasonDetails(urls <-chan string) <-chan Show {
	shows := make(chan Show)

	go func() {
		for u := range urls {
			fmt.Println(u)
			detailCollector := colly.NewCollector(colly.Async(true))

			detailCollector.OnHTML(".site-content", func(e *colly.HTMLElement) {
				title := e.ChildText(".entry-title")
				url := e.Request.URL.String()
				cover := e.ChildAttr("div.series-image img[src^='']", "src")
				description := e.ChildText(".series-desc")[16:]
				shows <- Show{Title: title,
					URL:         url,
					Cover:       cover,
					Description: description,
				}
				fmt.Println(e.Request.URL.String())
			})

			detailCollector.Visit(u)
		}
	}()

	return shows
}

func toJSON(s []Show) string {
	jsonBytes, err := json.Marshal(s)

	if err != nil {
		panic(err)
	}

	return string(jsonBytes)
}
