-- name: AddOauthAccount :one
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
RETURNING *;