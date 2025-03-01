package models

type SharedList struct {
	ID               uint `json:"id"`
	ListId           uint `json:"list_id"`
	SharedWithUserId uint `json:"shared_with_user_id"`
}
