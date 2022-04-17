-- +goose Up
-- +goose StatementBegin
ALTER TABLE todo_list MODIFY completed  TINYINT(1) unsigned DEFAULT '0';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE todo_list MODIFY completed   bit(1) NOT NULL DEFAULT false;
-- +goose StatementEnd
