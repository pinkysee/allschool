-- +goose Up
-- +goose StatementBegin
CREATE TABLE news(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    text VARCHAR(50000) NOT NULL,
    preview VARCHAR(255) NOT NULL,
    created_at VARCHAR(255) NOT NULL
);
-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS news;
-- +goose StatementEnd