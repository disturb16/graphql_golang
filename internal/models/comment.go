package models

type Comment struct{
	ID int64 `json:"id"`
	Name string `json:"name"`
	Content string `json:"content"`
	PostID int64 `json:"post_id"`
}