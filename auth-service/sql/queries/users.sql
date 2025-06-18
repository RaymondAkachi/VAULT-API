-- name: CreateUser :one
INSERT INTO users (id, username, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: AuthenticateUser :one
SELECT * FROM users
WHERE email=$1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id=$1;