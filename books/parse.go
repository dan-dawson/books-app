package books

import (
	"books-app/shared"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var BookList shared.Books

func BookParse(s string) (map[string]shared.Book, error) {

	fileContent, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
		return map[string]shared.Book{}, err
	}

	byteResult, _ := ioutil.ReadAll(fileContent)

	var newBooks shared.Books

	err = json.Unmarshal(byteResult, &newBooks)

	if err != nil {
		return map[string]shared.Book{}, err
	}

	BookList = newBooks
	booksMap := make(map[string]shared.Book)

	for _, v := range newBooks.AllBooks {
		thisBook := v
		booksMap[thisBook.Title] = thisBook
	}

	return booksMap, err
}

func UpdateJson(bookList map[string]shared.Book) string {
	_, err := os.Stat("books.json")
	if os.IsNotExist(err) {
		_, err := os.Create("books.json")
		if err != nil {
			return ""
		}
	}

	this, _ := json.MarshalIndent(bookList, "", "  ")

	err = ioutil.WriteFile("books.json", this, 0644)
	if err != nil {
		return ""
	}
	return string(this)
}
