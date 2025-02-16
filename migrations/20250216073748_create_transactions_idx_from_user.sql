-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS transactions_from_user_idx ON transactions USING HASH (from_user);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS transactions_from_user_idx;
-- +goose StatementEnd
