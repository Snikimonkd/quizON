LOCAL_DB_DSN:=$(shell grep -A1 'database' local_config.yaml | tail -n1 | sed "s/.*dsn: //g" | sed "s/\"//g")
APP_PORT:=$(shell grep -A1 'server' local_config.yaml | tail -n1 | sed "s/.*port: //g" | sed "s/\"//g")

create-migration:
	goose -dir migrations create $(NAME) sql

migrate-up:
	goose -dir migrations postgres "$(LOCAL_DB_DSN)" up

migrate-test:
	goose -dir test_migrations postgres "$(LOCAL_DB_DSN)" up

migrate-down:
	goose -dir migrations postgres "$(LOCAL_DB_DSN)" down

run:
	go run cmd/main.go

jet:
	@PATH=$(LOCAL_BIN):$(PATH) jet -dsn $(LOCAL_DB_DSN) -path=./internal/model -schema=public

run-docker:
	docker compose up