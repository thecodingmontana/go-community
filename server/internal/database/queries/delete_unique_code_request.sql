-- name: DeleteUniqueCodeRequest :exec
DELETE FROM "verification_codes"
WHERE email = $1;