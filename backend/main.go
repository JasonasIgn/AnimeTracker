// test
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type show struct {
	Title string
	URL   string
}

func main() {

	//Handling redirections
	http.HandleFunc("/current-season", current_season)

	log.Fatal(http.ListenAndServe(":1337", nil))
}

func current_season(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, get_current_season())
}

func get_current_season() string {
	shows := []show{}

	c := colly.NewCollector(
		colly.AllowedDomains("horriblesubs.info"),
	)

	c.OnHTML(".ind-show", func(e *colly.HTMLElement) {
		temp := show{}
		temp.Title = e.ChildText("a[title]")
		temp.URL = e.ChildAttr("a[href]", "href")
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

	json_bytes, err := json.Marshal(shows)

	if err != nil {
		panic(err)
	}

	return string(json_bytes)
}
