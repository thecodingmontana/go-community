// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: update_unique-code.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const updateUniqueCode = `-- name: UpdateUniqueCode :one
UPDATE "verification_codes"
SET code = $1,
    expires_at = $2,
    updated_at = $3
WHERE email = $4
RETURNING code
`

type UpdateUniqueCodeParams struct {
	Code      string
	ExpiresAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamp
	Email     string
}

func (q *Queries) UpdateUniqueCode(ctx context.Context, arg UpdateUniqueCodeParams) (string, error) {
	row := q.db.QueryRow(ctx, updateUniqueCode,
		arg.Code,
		arg.ExpiresAt,
		arg.UpdatedAt,
		arg.Email,
	)
	var code string
	err := row.Scan(&code)
	return code, err
}
