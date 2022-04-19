-- +goose Up
-- +goose StatementBegin
ALTER TABLE users MODIFY password VARCHAR(60)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users MODIFY password VARCHAR(30)
-- +goose StatementEnd
