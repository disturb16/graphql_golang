package models

type Post struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID string `json:"author_id"`
}
