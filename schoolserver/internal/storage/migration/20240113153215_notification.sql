-- +goose Up
-- +goose StatementBegin
create table notification(
   Text varchar(500),
   Class VARCHAR(255) NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notification;
-- +goose StatementEnd
