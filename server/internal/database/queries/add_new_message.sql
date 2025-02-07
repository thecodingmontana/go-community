-- name: AddNewMessage :exec
INSERT INTO "messages"(id, user_id, content, image_url)
VALUES($1, $2, $3, $4);