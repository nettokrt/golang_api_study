package store

import (
	"context"
	"database/sql"
	"errors"
)

var ( ErrPostNotFound = errors.New("post not found") 
)
type Storage struct {
	Posts interface{
		GetByID(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
	}

	Users interface{
		Create(context.Context, *User) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &UserStore{db},
	}
}