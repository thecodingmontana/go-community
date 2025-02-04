-- name: CreateUser :one
INSERT INTO "users"(id, email, username, email_verified, avatar)
VALUES($1, $2, $3, $4, $5)
RETURNING *;