-- name: GenerateUniqueCode :one
INSERT INTO "verification_codes"(id, email, code, expires_at)
VALUES($1, $2, $3, $4)
RETURNING *;