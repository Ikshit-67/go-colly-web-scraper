package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// type product struct {
// 	Name     string `json: "name"`
// 	Price    string `json: "price"`
// 	ImageUrl string `json: "image_url"`
// }

func main() {

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
		// fmt.Println("Author:\n", h.Text)
		div := h.DOM

		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		fmt.Printf("\nQuote: %v\n   - by %v\n\n\n", quote, author)
	})

	c.Visit("http://quotes.toscrape.com/")

}
