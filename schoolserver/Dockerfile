FROM golang

WORKDIR /schoolserver
COPY . .

RUN go mod tidy
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD sh -c 'cd internal/storage/migration && goose postgres "host=db user=postgres database=schooldb password=postgres sslmode=disable" up && cd ../../../cmd/apiserver && go run main.go'
