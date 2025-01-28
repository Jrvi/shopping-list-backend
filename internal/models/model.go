package models

type Product struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	CategoryID string `json:"category_id"`
	ListID     string `json:"list_id"`
}

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type List struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
