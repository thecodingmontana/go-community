-- name: UpdateUniqueCode :one
UPDATE "verification_codes"
SET code = $1,
    expires_at = $2,
    updated_at = $3
WHERE email = $4
RETURNING code;