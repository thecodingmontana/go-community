// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: add_oauth_account.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addOauthAccount = `-- name: AddOauthAccount :one
INSERT INTO "oauth_account"(
        id,
        user_id,
        provider,
        provider_user_id,
        access_token,
        refresh_token,
        expires_at
    )
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, provider, provider_user_id, access_token, refresh_token, expires_at, created_at, updated_at
`

type AddOauthAccountParams struct {
	ID             string
	UserID         string
	Provider       string
	ProviderUserID string
	AccessToken    string
	RefreshToken   pgtype.Text
	ExpiresAt      pgtype.Timestamptz
}

func (q *Queries) AddOauthAccount(ctx context.Context, arg AddOauthAccountParams) (OauthAccount, error) {
	row := q.db.QueryRow(ctx, addOauthAccount,
		arg.ID,
		arg.UserID,
		arg.Provider,
		arg.ProviderUserID,
		arg.AccessToken,
		arg.RefreshToken,
		arg.ExpiresAt,
	)
	var i OauthAccount
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Provider,
		&i.ProviderUserID,
		&i.AccessToken,
		&i.RefreshToken,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
