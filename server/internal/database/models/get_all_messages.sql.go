// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: get_all_messages.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllMessages = `-- name: GetAllMessages :many
SELECT 
    message.id,
    message.content,
    message.deleted,
    message.image_url,
    message.file_url,
    message.created_at,
    message.updated_at,
    u.id AS by_id,
    u.email AS by_email,
    u.username AS by_username,
    u.avatar AS by_avatar,
    u.email_verified AS by_email_verified
FROM "messages" message
JOIN "users" u ON message.user_id = u.id
ORDER BY message.created_at DESC
`

type GetAllMessagesRow struct {
	ID              string
	Content         string
	Deleted         bool
	ImageUrl        string
	FileUrl         string
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
	ByID            string
	ByEmail         string
	ByUsername      string
	ByAvatar        pgtype.Text
	ByEmailVerified bool
}

func (q *Queries) GetAllMessages(ctx context.Context) ([]GetAllMessagesRow, error) {
	rows, err := q.db.Query(ctx, getAllMessages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllMessagesRow
	for rows.Next() {
		var i GetAllMessagesRow
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.Deleted,
			&i.ImageUrl,
			&i.FileUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ByID,
			&i.ByEmail,
			&i.ByUsername,
			&i.ByAvatar,
			&i.ByEmailVerified,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
