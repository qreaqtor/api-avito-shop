-- +goose Up
-- +goose StatementBegin
INSERT INTO merch ("merch_type", "price") VALUES
('t-shirt', 80),
('cup', 20),
('book', 50),
('pen', 10),
('powerbank', 200),
('hoody', 300),
('umbrella', 200),
('socks', 10),
('wallet', 50),
('pink-hoody', 500);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM merch 
WHERE "merch_type" IN ('t-shirt', 'cup', 'book', 'pen', 'powerbank', 'hoody', 'umbrella', 'socks', 'wallet', 'pink-hoody');
-- +goose StatementEnd
