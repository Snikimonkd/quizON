-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id serial PRIMARY KEY NOT NULL,
    login text NOT NULL UNIQUE,
    password bytea NOT NULL
);

CREATE TABLE cookies(
    user_id integer NOT NULL,
    value uuid NOT NULL DEFAULT gen_random_uuid(),
    expires_at timestamptz NOT NULL DEFAULT now() + interval '1 day',
    FOREIGN KEY(user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cookies;
DROP TABLE users;
-- +goose StatementEnd
