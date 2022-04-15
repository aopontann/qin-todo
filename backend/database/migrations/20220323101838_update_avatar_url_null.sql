-- +goose Up
-- +goose StatementBegin
ALTER TABLE users MODIFY avatar_url VARCHAR(255)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users MODIFY avatar_url VARCHAR(255) NOT NULL
-- +goose StatementEnd
