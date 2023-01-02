-- +goose Up
-- +goose StatementBegin
ALTER TABLE games ADD COLUMN registered_teams integer NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE games DROP COLUMN registered_teams;
-- +goose StatementEnd
