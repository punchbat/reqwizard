-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_gender AS ENUM ('male', 'female', 'other');

ALTER TABLE users
    ALTER COLUMN gender TYPE user_gender USING gender::user_gender;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    ALTER COLUMN gender TYPE VARCHAR(6) USING gender::VARCHAR;

DROP TYPE user_gender;
-- +goose StatementEnd