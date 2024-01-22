-- +goose Up
-- +goose StatementBegin
create table schedule(
    classname varchar(5),
    lessons varchar(500)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS schedule;
-- +goose StatementEnd
