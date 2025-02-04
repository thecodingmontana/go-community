-- name: FindUserUniqueCode :one
SELECT *
FROM "verification_codes"
WHERE email = $1;