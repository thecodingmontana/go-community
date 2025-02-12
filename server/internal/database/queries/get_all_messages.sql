-- name: GetAllMessages :many
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
ORDER BY message.created_at DESC;