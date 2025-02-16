-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS transactions_to_user_idx ON transactions USING HASH (to_user);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS transactions_to_user_idx;
-- +goose StatementEnd
