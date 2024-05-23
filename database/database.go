package database

import (
	"awesomeProject/models"
	"database/sql"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=Log680968amr dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetCommentsByPostID(postId int) ([]*models.Comment, error) {
	rows, err := db.Query("SELECT id, post_id, parent_id, content, author_id, created_at FROM comments WHERE post_id = $1", postId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.AuthorID, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func CreateComment(comment models.Comment) error {
	_, err := db.Exec("INSERT INTO comments (post_id, parent_id, content, author_id, created_at) VALUES ($1, $2, $3, $4, NOW())", comment.PostID, comment.ParentID, comment.Content, comment.AuthorID)
	return err
}

func GetLatestCommentByPostID(postId int) (*models.Comment, error) {
	var comment models.Comment
	err := db.QueryRow("SELECT id, post_id, parent_id, content, author_id, created_at FROM comments WHERE post_id = $1 ORDER BY created_at DESC LIMIT 1", postId).Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.Content, &comment.AuthorID, &comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
func GetPostByID(id int) (*models.Post, error) {
	var post models.Post
	err := db.QueryRow("SELECT id, title, content, allow_comments, created_at FROM posts WHERE id = $1", id).Scan(&post.ID, &post.Title, &post.Content, &post.AllowComments, &post.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func CreatePost(post models.Post) *models.Post {
	// Устанавливаем текущую дату и время перед вставкой
	post.CreatedAt = time.Now()

	// Выполняем SQL-запрос для вставки новой записи в таблицу posts
	row := db.QueryRow("INSERT INTO posts (title, content, allow_comments, created_at) VALUES ($1, $2, $3, $4) RETURNING id", post.Title, post.Content, post.AllowComments, post.CreatedAt)

	// Получаем идентификатор вставленного поста
	err := row.Scan(&post.ID)
	if err != nil {
		return nil
	}

	return &post
}
