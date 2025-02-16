-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    "username" varchar(32) PRIMARY KEY,
    "password" varchar(255) NOT NULL,
    "coins" integer CHECK (coins >= 0) DEFAULT 1000,
    "created_at" TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
