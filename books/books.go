package books

var Books = []Book{
	sampleBook,
}

type Book struct {
	Title string `json:"title"`
	Pages Pages  `json:"pages"`
}

var sampleBook = Book{
	Title: "An amazing book",
	Pages: []Page{
		{
			Text:  "This is the start of a amazing book",
			Image: "th.jpeg",
		}, {
			Text:  "This is the next page of a amazing book",
			Image: "",
		}, {
			Text:  "This is the third page of a amazing book",
			Image: "",
		}, {
			Text:  "This is the fourth page of a amazing book",
			Image: "",
		},
	},
}
