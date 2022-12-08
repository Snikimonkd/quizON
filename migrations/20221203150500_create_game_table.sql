-- +goose Up
-- +goose StatementBegin
CREATE TABLE games(
    id SERIAL PRIMARY KEY NOT NULL,
    name text NOT NULL,
    description text NOT NULL,
    date timestamptz NOT NULL,
    teams_amount integer NOT NULL,
    price_per_person integer NOT NULL,
    location text NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL,
    updated_at timestamptz DEFAULT now() NOT NULL,
    created_by integer NOT NULL,
    FOREIGN KEY(created_by) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE games;
-- +goose StatementEnd
