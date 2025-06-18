package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	db "github.com/RaymondAkachi/VAULT-API/auth-service/internal/database"
	utils "github.com/RaymondAkachi/VAULT-API/auth-service/utils"
	"github.com/google/uuid"
)

type User struct {
    ID       int64
    Username string
    Email    string
    Password string
}

func CreateUser(ctx context.Context, username, email, password string) (db.User,error) {
    // Hash password
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        return db.User{}, err
    }
    // query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
    // _, err = apiConfig.DB.Exec(query, username, email, string(hashedPassword))
    params := db.CreateUserParams{
        ID: uuid.New(),
        Username: username,
        Email:  email,
        Password: hashedPassword,
    }

    user, err := queries.CreateUser(ctx, params)

    if err != nil {
        log.Fatalf("Could not create user: %v", err)
        return db.User{}, err
    }
    
    fmt.Println("User")

    return user, err
}

func AuthenticateUser(ctx context.Context, email, password string) (db.User, error) {
    // query := `SELECT password FROM users WHERE email=$1`
    // err := conn.QueryRow(query, email).Scan(&hashedPassword)
    // if err != nil {
    //     return errors.New("user not found")
    user, err:= queries.AuthenticateUser(ctx, email)
    if err != nil {
        return db.User{}, errors.New("user not found")
    }
	is_valid := utils.CheckPasswordHash(password, user.Password) 
	if !is_valid {
		return db.User{}, errors.New("invalid user password entered")
	}
    return user, nil
}
