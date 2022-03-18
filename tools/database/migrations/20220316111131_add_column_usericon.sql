-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD avatar_url VARCHAR(255) NOT NULL AFTER token
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP avatar_url
-- +goose StatementEnd
