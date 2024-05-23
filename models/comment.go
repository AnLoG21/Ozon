package models

type Comment struct {
	ID        int
	PostID    int
	ParentID  *int
	Content   string
	AuthorID  int
	CreatedAt string
}
