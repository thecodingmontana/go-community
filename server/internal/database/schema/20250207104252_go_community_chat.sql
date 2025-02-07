-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "messages" (
	"id" text PRIMARY KEY NOT NULL,
	"user_id" text NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "content" text NOT NULL,
    "image_url" varchar(255) NOT NULL,
    "deleted" boolean NOT NULL DEFAULT false,
	"created_at" timestamp (3) DEFAULT now() NOT NULL,
	"updated_at" timestamp (3) DEFAULT now() NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "messages";
-- +goose StatementEnd
