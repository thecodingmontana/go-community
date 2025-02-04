-- name: FindUserOauthAccount :one
SELECT *
FROM "oauth_account"
WHERE user_id = $1;