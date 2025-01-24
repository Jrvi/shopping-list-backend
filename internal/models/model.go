package models

type Product struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	CategoryID string `json:"category_id"`
}

type Category struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
