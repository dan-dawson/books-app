package main

import (
	"books-app/books"
	"books-app/shared"
	"fyne.io/fyne/app"
)

var allBooks map[string]shared.Book

func main() {
	a := app.New()
	w := a.NewWindow("book")

	allBooks, _ = books.BookParse("books/books.json")

	u := shared.UI{}

	u.LoadUI(w, allBooks, a)
	a.Run()
	books.UpdateJson(allBooks)
}
