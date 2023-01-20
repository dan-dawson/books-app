package books

type Page struct {
	Image string `json:"image"`
	Text  string `json:"text"`
}

type Pages []Page
