-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS inventory_items (
    username varchar(32) references users(name),
    merch_name varchar(255) references merch(name),
    quantity integer NOT NULL,
    PRIMARY KEY (username, merch_name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS inventory_items; 
-- +goose StatementEnd
