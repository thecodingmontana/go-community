-- name: FindUserByEmail :one
SELECT *
FROM "users"
WHERE email = $1;