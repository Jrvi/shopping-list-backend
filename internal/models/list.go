package models

type List struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	UserId   uint   `json:"user_id"`
	IsShared bool   `json:"is_shared"`
}
