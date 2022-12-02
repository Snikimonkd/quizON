create-migration:
	goose -dir migrations create $(NAME) sql

migrate-up:
	goose -dir migrations postgres "$(LOCAL_DB_DSN)" up

migrate-down:
	goose -dir migrations postgres "$(LOCAL_DB_DSN)" down