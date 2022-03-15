-- +goose Up
-- +goose StatementBegin
CREATE TABLE todo_list(
  id VARCHAR(26) PRIMARY KEY NOT NULL,
  content VARCHAR(255) NOT NULL,
  completed bit(1) NOT NULL DEFAULT false,
  execution_date DATETIME,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  user_id VARCHAR(26) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS todo_list;
-- +goose StatementEnd
