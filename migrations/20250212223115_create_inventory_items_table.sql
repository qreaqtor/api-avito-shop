-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS inventory_items (
    "username" varchar(32) references users("username"),
    "merch_type" varchar(255) references merch("merch_type")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory_items; 
-- +goose StatementEnd
