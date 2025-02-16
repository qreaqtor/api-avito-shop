-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS merch_type_idx ON merch USING HASH ("merch_type");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS merch_type_idx;
-- +goose StatementEnd
