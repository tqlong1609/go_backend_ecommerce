APP_NAME = "go-web-app"

# Goose default values
GOOSE_DRIVER ?= postgres
GOOSE_DBSTRING = "host=localhost port=5432 user=postgres password=quanglong dbname=store_db sslmode=disable"
GOOSE_MIGRATIONS_DIR ?= sql/schema

run:
	go run ./cmd/server/

# Section I: docker commands

docker-build:
	docker-compose up -d --build
docker-up:
	docker-compose up -d
docker-down:
	docker-compose down
docker-stop:
	docker-compose stop

# force stop all containers
# use when:
# - docker-compose stop not working
# - emergency stop
docker-kill:
	docker-compose kill

# remove all containers, images, volumes, and networks
# use when: 
# - clear all data in application, database, ...
# - image outdated
# - note: this command will remove all data in database
docker-clean:
	docker-compose down --rmi all -v

.PHONY: run docker-build docker-up docker-down docker-stop docker-kill docker-clean

### End: docker ###

# Section II: goose commands

create_migration:
	goose -dir=$(GOOSE_MIGRATIONS_DIR) create $(name) sql
up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_DIR) up
down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_DIR) down
reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_DIR) reset
up-by-one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATIONS_DIR) up-by-one
# Specifies that the targets (run, up, down, reset) are not actual files or directories, but commands to run.
# Avoid Make mistaking the target for a file if there is a file of the same name in the directory
.PHONY: run up down reset sqlgen

# End: goose commands

# Section III: sqlc commands

