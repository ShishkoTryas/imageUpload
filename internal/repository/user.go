package repository

import (
	"context"
	"database/sql"
	"imageUpload/internal/domain"
)

type user struct {
	db *sql.DB
}

func New(db *sql.DB) *user {
	return &user{
		db: db,
	}
}

func (u *user) Create(ctx context.Context, user domain.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, password) values ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}

func (u *user) GetUser(ctx context.Context, email, password string) (domain.User, error) {
	var user domain.User
	err := u.db.QueryRow("SELECT id, name, email FROM users WHERE email=$1 AND password=$2", email, password).
		Scan(&user.Id, &user.Name, &user.Email)

	return user, err
}
