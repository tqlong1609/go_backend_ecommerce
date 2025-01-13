APP_NAME = "go-web-app"

# Goose default values
GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING = "host=localhost port=5432 user=postgres password=quanglong dbname=store_db sslmode=disable"
GOOSE_MIGRATIONS_DIR ?= sql/schema

run:
	go run ./cmd/server/

# goose commands

up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_DIR) up
down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_DIR) down
reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_DIR) reset

# Specifies that the targets (run, up, down, reset) are not actual files or directories, but commands to run.
# Avoid Make mistaking the target for a file if there is a file of the same name in the directory
.PHONY: run up down reset