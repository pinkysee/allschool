-- +goose Up
-- +goose StatementBegin
create table test(
    id SERIAL PRIMARY KEY,
    name varchar(55),
    lesson varchar(50),
    quetion VARCHAR(500000),
    teachername varchar(100),
    komu varchar(500),
    deadline varchar(50),
    answer VARCHAR(500000),
    rightans VARCHAR(5000)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS test;
-- +goose StatementEnd