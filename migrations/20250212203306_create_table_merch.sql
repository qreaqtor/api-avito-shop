-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS merch (
    "merch_type" varchar(255) PRIMARY KEY,
    "price" integer NOT NULL,
    "created_at" TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS merch;
-- +goose StatementEnd
