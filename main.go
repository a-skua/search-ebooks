package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"search-ebooks/search"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, heroku"))
	})
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		word := r.FormValue("word")
		f := func(w string, s search.Search, c chan<- []search.Book) {
			res, err := http.Get(s.URL())
			if err != nil {
				log.Println(err)
				c <- nil
				return
			}
			defer res.Body.Close()
			books, err := s.Scrape(res.Body, w)
			if err != nil {
				log.Println(err)
				c <- nil
				return
			}
			c <- books
		}

		stores := search.GetStores()
		l := len(stores)
		ch := make(chan []search.Book)
		for i := 0; i < l; i++ {
			go f(word, stores[i], ch)
		}

		var books []search.Book
		for i := 0; i < l; i++ {
			books = append(books, <-ch...)
		}
		w.Write([]byte(fmt.Sprint(books)))
	})

	log.Println("ListenAndServe:", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
