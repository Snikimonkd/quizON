-- +goose Up
-- +goose StatementBegin
INSERT INTO users (login, password) VALUES('test-user', '$2a$10$uhaUyLpOL8WzAxa85fzCx.qTqjvBLLpwosspveypJ/g7UZntFr7KW');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM users WHERE login = 'test-user';
-- +goose StatementEnd
