-- +goose Up
-- +goose StatementBegin
CREATE TABLE testdb(id integer);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE testdb;
-- +goose StatementEnd
