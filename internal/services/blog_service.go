package services

import (
	"github.com/disturb16/graphql_golang/internal/models"
)

// Posts returns list of posts
func (s *Service) Posts() ([]models.Post, error) {

	posts := []models.Post{}
	qry := "select id, title, content, author_id from POSTS"

	rows, err := s.db.Query(qry)

	if err != nil {
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		post := models.Post{}

		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)

		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	return posts, err
}

// Comments returns list of comments
func (s *Service) Comments() ([]models.Comment, error) {
	comments := []models.Comment{}
	qry := "select id, name, content, post_id from COMMENTS"

	rows, err := s.db.Query(qry)

	if err != nil {
		return comments, err
	}

	defer rows.Close()

	for rows.Next() {
		comment := models.Comment{}

		err = rows.Scan(&comment.ID, &comment.Name, &comment.Content, &comment.PostID)

		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, err
}

// CommentsByPost returns list of comments by post
func (s *Service) CommentsByPost(postID int64) ([]models.Comment, error) {

	comments := []models.Comment{}
	qry := "select id, name, content, post_id from COMMENTS where post_id = ?"

	rows, err := s.db.Query(qry, postID)

	if err != nil {
		return comments, err
	}

	defer rows.Close()

	for rows.Next() {
		comment := models.Comment{}

		err = rows.Scan(&comment.ID, &comment.Name, &comment.Content, &comment.PostID)

		if err != nil {
			return comments, err
		}

		comments = append(comments, comment)
	}

	return comments, err
}

// AuthorByID returns specific author
func (s *Service) AuthorByID(authorID int) (models.Author, error) {
	author := models.Author{}
	qry := "select id, name, email from AUTHORS where id = ?"
	rows, err := s.db.Query(qry, authorID)

	if err != nil {
		return author, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&author.ID, &author.Name, &author.Email)

		if err != nil {
			return author, err
		}
	}

	return author, err
}

// PostsByAuthor returns list of posts by author
func (s *Service) PostsByAuthor(authorID int64) ([]models.Post, error) {
	posts := []models.Post{}
	qry := "select id, title, content, author_id from POSTS where author_id = ?"

	rows, err := s.db.Query(qry, authorID)

	if err != nil {
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		post := models.Post{}

		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)

		if err != nil {
			return posts, err
		}

		posts = append(posts, post)
	}

	return posts, err
}

// AddPost creates a new post in database
func (s *Service) AddPost(title, content string, authorID int) (int64, error) {
	qry := "insert into POSTS (title, content, author_id) values(?, ?, ?)"

	result, err := s.db.Exec(qry, title, content, authorID)

	if err != nil {
		return 0, err
	}

	insertID, _ := result.LastInsertId()

	return insertID, nil
}

// PostByID returns specific  post
func (s *Service) PostByID(id int64) (models.Post, error) {
	post := models.Post{}

	qry := "select id, title, content, author_id from POSTS where id = ?"
	result := s.db.QueryRow(qry, id)
	err := result.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID)

	if err != nil {
		return post, err
	}

	return post, nil
}

// AddComment creates a new comment in database
func (s *Service) AddComment(name, content string, postID int) (int64, error) {

	qry := "insert into COMMENTS (name, content, post_id) values(?, ?, ?)"

	result, err := s.db.Exec(qry, name, content, postID)

	if err != nil {
		return 0, err
	}

	insertID, _ := result.LastInsertId()

	return insertID, nil
}

// CommentByID returns specific comment
func (s *Service) CommentByID(commentID int64) (models.Comment, error) {

	comment := models.Comment{}

	qry := "select id, name, content, post_id from COMMENTS where id = ?"
	result := s.db.QueryRow(qry, commentID)
	err := result.Scan(&comment.ID, &comment.Name, &comment.Content, &comment.PostID)

	if err != nil {
		return comment, err
	}

	return comment, nil
}
