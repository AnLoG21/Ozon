package models

import "time"

type Post struct {
	ID            int
	Title         string
	Content       string
	AuthorID      int
	AllowComments bool
	CreatedAt     time.Time
}

func (p Post) Error() string {
	//TODO implement me
	panic("implement me")
}
