-- +goose Up
-- +goose StatementBegin
create table docs (
    name varchar(255) not null,
    path varchar(255) not null,
    page varchar(300)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS docs;
-- +goose StatementEnd
