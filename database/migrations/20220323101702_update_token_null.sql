-- +goose Up
-- +goose StatementBegin
ALTER TABLE users MODIFY token VARCHAR(255) UNIQUE
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users MODIFY token VARCHAR(255) NOT NULL UNIQUE
-- +goose StatementEnd
