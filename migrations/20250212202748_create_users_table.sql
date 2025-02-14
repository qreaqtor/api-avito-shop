-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    name varchar(32) PRIMARY KEY,
    password varchar(255) NOT NULL,
    coins integer CHECK (coins >= 0) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
