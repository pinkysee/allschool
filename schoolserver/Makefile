.PHONY: build
containername = schooldb
dbname = schooldb
build:
	go build -v ./cmd/apiserver
run:
	cd cmd/apiserver && \
	go run main.go
migration:
	cd internal/storage/migration && \
	goose postgres "host=localhost user=postgres database=schooldb password=postgres sslmode=disable" up
migrationdown: 
	cd internal/storage/migration && \
	goose postgres "host=localhost user=postgres database=schooldb password=postgres sslmode=disable" down-to 0
dockerdb:
	docker-compose -f docker-compose.dev.yaml up -d
testhandler:
	go test ./...

.DEFAULT_GOAL := run