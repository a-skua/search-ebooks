package search

import (
	"github.com/PuerkitoBio/goquery"
	"io"
	"regexp"
)

var Cap = 100

type Book struct {
	ISBN, Title, ImageURL, Price string
}
type Search interface {
	Scrape(io.Reader, string) ([]Book, error)
}

type OreillyJP struct{}

func (OreillyJP) Scrape(data io.Reader, word string) ([]Book, error) {
	doc, err := goquery.NewDocumentFromReader(data)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(word)
	books := make([]Book, 0, Cap)
	doc.Find("#bookTable").Find("tr").Each(func(i int, s *goquery.Selection) {
		book := Book{
			ISBN:  s.Find(".isbn").Text(),
			Title: s.Find(".title").Text(),
			Price: s.Find(".price").Text() + " JPY",
		}
		book.ImageURL = "//oreilly.co.jp/books/images/picture_small" + book.ISBN + ".gif"

		if r.MatchString(book.Title) {
			books = append(books, book)
		}
	})
	return books, nil
}
