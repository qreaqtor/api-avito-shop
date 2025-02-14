-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS merch (
    name varchar(255) PRIMARY KEY,
    price integer NOT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS merch;
-- +goose StatementEnd
