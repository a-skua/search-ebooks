package search

import (
	"strings"
	"testing"
)

const html = `
<table id="bookTable">
<tbody>
<tr>
  <td class="isbn">999-8-77777-666-55</td>
  <td class="title"><a href="foo">XML解析</a></td>
  <td class="price">1,234</td>
</tr>
<tr>
  <td class="isbn">888-7-66666-555-44</td>
  <td class="title"><a href="foo">JSON解析</a></td>
  <td class="price">2,345</td>
</tr>
<tr>
  <td class="isbn">777-6-55555-444-33</td>
  <td class="title"><a href="foo">HTML解析</a></td>
  <td class="price">3,456</td>
</tr>
<tr>
  <td class="isbn">666-5-44444-333-22</td>
  <td class="title"><a href="foo">詳解Linux</a></td>
  <td class="price">4,567</td>
</tr>
</tbody>
</table>
`

func TestOreillySearch(t *testing.T) {
	var store OreillyJP
	data := strings.NewReader(html)
	books, err := store.Scrape(data, "HTML")
	if err != nil {
		t.Error(err)
		return
	}
	if l := len(books); l < 1 || l > 1 {
		t.Error("length:", l, books)
		return
	}

	book := books[0]
	if isbn := "777-6-55555-444-33"; book.ISBN != isbn {
		t.Error(book.ISBN, "!=", isbn)
	}

	if title := "HTML解析"; book.Title != title {
		t.Error(book.Title, "!=", title)
	}

	if imgURL := "//oreilly.co.jp/books/images/picture_small777-6-55555-444-33.gif"; book.ImageURL != imgURL {
		t.Error(book.ImageURL, "!=", imgURL)
	}

	if price := "3,456 JPY"; book.Price != price {
		t.Error(book.Price, "!=", price)
	}
}

func TestOreillySearchCheckNum(t *testing.T) {
	var store OreillyJP

	data := strings.NewReader(html)
	books, err := store.Scrape(data, "HTML")
	if err != nil {
		t.Error(err)
		return
	}
	if l := len(books); l != 1 {
		t.Error("length is not 1:", l)
	}

	data = strings.NewReader(html)
	books, err = store.Scrape(data, "ML")
	if err != nil {
		t.Error(err)
		return
	}
	if l := len(books); l != 2 {
		t.Error("length is not 2:", l)
	}

	data = strings.NewReader(html)
	books, err = store.Scrape(data, "解析")
	if err != nil {
		t.Error(err)
		return
	}
	if l := len(books); l != 3 {
		t.Error("length is not 3:", l)
	}

	data = strings.NewReader(html)
	books, err = store.Scrape(data, "")
	if err != nil {
		t.Error(err)
		return
	}
	if l := len(books); l != 4 {
		t.Error("length is not 4:", l)
	}
}
