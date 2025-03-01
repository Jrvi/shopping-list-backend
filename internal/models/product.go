package models

type Product struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	CategoryID uint   `json:"category_id"`
	ListID     uint   `json:"list_id"`
}
