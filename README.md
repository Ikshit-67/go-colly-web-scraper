# go-colly-web-scraper

## Description

The go-colly-web-scraper project is a Go-based web scraping application using the Colly framework. It demonstrates the capabilities of Colly for efficiently crawling and extracting data from web pages. The application is designed to scrape quotes from a sample website, showcasing the use of Colly's features such as concurrent crawling, HTML element selection, and JSON data export. Read more about colly [here](https://go-colly.org/)

## Features

- Concurrent and asynchronous web scraping with Colly.
- Selector-based extraction of data from HTML elements.
- JSON export of scraped data for further analysis.

## Installation

### Clone the Repository

```bash 
git clone https://github.com/Ikshit-67/go-colly-web-scraper.git
cd go-colly-web-scraper
```

### Install Dependencies

```bash
go get -u github.com/gocolly/colly/v2
```

### Run the code
- Make sure you have go installed on your system, check it by running the command below:
```bash
go version
```
- If you do not have go installed, install it by clicking [here](https://go.dev/doc/install)

- Once go is properly installed then run the main.go file by:
```bash
go run main.go
```
- Make sure you are in the correct directory

> You can also delete the allQuotes.json file. This file will be automatically generated when you run main.go

### Useful links
- [Go docs](https://go.dev/doc/)
- [Colly docs](https://go-colly.org/docs/)

Made with ❤️, by Ikshit C.

