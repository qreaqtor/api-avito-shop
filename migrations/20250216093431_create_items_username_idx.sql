-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS items_username_idx ON inventory_items USING HASH (username)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS items_username_idx;
-- +goose StatementEnd
