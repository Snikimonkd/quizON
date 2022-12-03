LOCAL_DB_DSN:=$(shell grep -A1 'database' config.yaml | tail -n1 | sed "s/.*dsn: //g" | sed "s/\"//g")

create-migration:
	goose -dir migrations create $(NAME) sql

migrate-up:
	goose -dir migrations postgres "$(LOCAL_DB_DSN)" up

migrate-down:
	goose -dir migrations postgres "$(LOCAL_DB_DSN)" down

run:
	go run cmd/main.go

jet:
	@PATH=$(LOCAL_BIN):$(PATH) jet -dsn $(LOCAL_DB_DSN) -path=./internal/temp/model -schema=public