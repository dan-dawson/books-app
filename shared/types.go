package shared

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type UI struct {
	CurrentPage *Page
	Pages       Pages
	Content     *fyne.Container
	Tools       Toolbar
}

type Toolbar struct {
	Plus  *widget.Toolbar
	Minus *widget.Toolbar
}

type Page struct {
	Image string `json:"image"`
	Text  string `json:"text"`
}

type Pages []Page

type Books struct {
	AllBooks []Book `json:"books"`
}

type Book struct {
	Title string `json:"title"`
	Pages Pages  `json:"pages"`
}
