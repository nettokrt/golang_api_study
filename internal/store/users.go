package store

import (
  "context"
  "database/sql"
)

type User struct {
	ID 				int64 	`json:"id"`
	Username 	string 	`json:"username"`
	Email 		string 	`json:"email"`
	Password 	string 	`json:"-"`
	Phone 		string 	`json:"phone"`	
	CreatedAt string 	`json:"created_at"`
}
type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
  INSERT INTO users (username, email, password, Phone)
  VALUES ($1, $2, $3) RETURNING id, created_at`
	err := s.db.QueryRowContext(
			ctx, 
			query, 
			user.Username, 
			user.Email, 
			user.Phone,   // added phone number field
			user.Password,
		).Scan(
			&user.ID, 
			&user.CreatedAt,
		)
			if err != nil {
				return err
			}
	
	return nil
}
