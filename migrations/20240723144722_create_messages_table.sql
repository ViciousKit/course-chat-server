-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    "from" BIGINT NOT NULL,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    chat_id BIGINT NOT NULL REFERENCES chats(id) ON DELETE CASCADE
)
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd