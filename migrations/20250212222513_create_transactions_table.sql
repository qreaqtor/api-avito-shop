-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
    "from_user" varchar(32) references users("username"),
    "to_user" varchar(32) references users("username"),
    "amount" integer NOT NULL,
    "created_at" TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
