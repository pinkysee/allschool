-- +goose Up
-- +goose StatementBegin
create table iframe(
   id SERIAL PRIMARY KEY,
   page VARCHAR(255) NOT NULL,
   path VARCHAR(255) NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS iframe;
-- +goose StatementEnd
