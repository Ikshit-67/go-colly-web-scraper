package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// type product struct {
// 	Name     string `json: "name"`
// 	Price    string `json: "price"`
// 	ImageUrl string `json: "image_url"`
// }

// * Quote struct
type Quote struct {
	Text   string `json:"quote_text"`
	Author string `json:"quote_author"`
}

func main() {

	// * Initializing an empty array of Quote struct for all quotes
	allQuotes := []Quote{}

	// * New colly instance
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	// * On Request callback
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// * On Error callback
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	// * On Response callback
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		fmt.Println("Response code", r.StatusCode)
	})

	// * On HTML callbacks
	// for .text
	// c.OnHTML(".text", func(h *colly.HTMLElement) {
	// 	fmt.Println("Quote:\n", h.Text)
	// })

	// for .author
	// c.OnHTML(".author", func(h *colly.HTMLElement) {
	// 	fmt.Println("Author:\n", h.Text)
	// })

	// for .quote
	// .quote contains .text and .author
	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		text := div.Find(".text").Text()
		author := div.Find(".author").Text()

		// Below line will print all the quotes and authors
		// fmt.Printf("\nQuote: %v\n   - by %v\n\n\n", quote, author)

		// * New instance of type Quote
		q := Quote{
			Text:   text,
			Author: author,
		}

		// * Append q in allQuotes array
		allQuotes = append(allQuotes, q)
	})

	c.Visit("http://quotes.toscrape.com/")

	// fmt.Println(allQuotes)
	data, err := json.Marshal(allQuotes)

	if err != nil {
		fmt.Println(err.Error())
	}

	os.WriteFile("allQuotes.json", data, 0644)

	fmt.Println("JSON data written to allQuotes.json successfully.")
}
