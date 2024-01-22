-- +goose Up
-- +goose StatementBegin
create table class(
   Name varchar(5)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS class;
-- +goose StatementEnd
