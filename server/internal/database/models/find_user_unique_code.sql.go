// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: find_user_unique_code.sql

package models

import (
	"context"
)

const findUserUniqueCode = `-- name: FindUserUniqueCode :one
SELECT id, code, email, expires_at, created_at, updated_at
FROM "verification_codes"
WHERE email = $1
`

func (q *Queries) FindUserUniqueCode(ctx context.Context, email string) (VerificationCode, error) {
	row := q.db.QueryRow(ctx, findUserUniqueCode, email)
	var i VerificationCode
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Email,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
