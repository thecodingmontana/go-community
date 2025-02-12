-- name: AddNewMessage :one
INSERT INTO "messages"(id, user_id, content, image_url, file_url)
VALUES($1, $2, $3, $4, $5)
RETURNING *;