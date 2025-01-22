-- +goose Up
-- +goose StatementBegin
ALTER TABLE messages
RENAME column "from" to author_id
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE messages
    RENAME column author_id to "from"
-- +goose StatementEnd