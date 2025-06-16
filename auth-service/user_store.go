package main

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	id int64 `db:"id"`
	Username string `db:"username"`
	Email string `db:"email"`
	Password string `db:"password_hash"`
}

func CreateUser(ctx context.Context, username, email, password string) error {
	// Hashes password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	
	query := `INSERT INTO user (username, email, password) VAULES ($1, $2, $3)`
	_, err = DB.Exec(query, username, email, string(hashedPassword))
	return err

}

func AuthenticateUser(email, password string) error {
	var hashedPassword string
	query := `SELECT password FROM users WHERE email=$1`
	err := DB.QueryRow(query, email).Scan(&hashedPassword)
	if err != nil {
		return errors.New("user nor found")
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// func FindUserByEmail(ctx context.Context, email string) (*User, error) {
// 	var user User
// 	err := DB.GetContext(ctx, &user, "SELECT * FROM users WHERE email=$1", email)
// 	if err != nil {
// 		return nil, errors.New("user not found")
// 	}
// 	return &user, nil
// }