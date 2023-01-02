-- +goose Up
-- +goose StatementBegin
CREATE TABLE registrations (
    id SERIAL PRIMARY KEY NOT NULL,
    game_id integer NOT NULL,
    team_name text NOT NULL,
    captain_name text NOT NULL,
    phone text NOT NULL,
    telega text NOT NULL,
    amount integer NOT NULL,
    registration_number integer NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    FOREIGN KEY(game_id) REFERENCES games(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE registrations;
-- +goose StatementEnd
