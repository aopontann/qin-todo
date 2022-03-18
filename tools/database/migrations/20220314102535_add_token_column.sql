-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD token VARCHAR(255) NOT NULL UNIQUE AFTER password
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP token
-- +goose StatementEnd
