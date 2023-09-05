-- +goose Up
-- +goose StatementBegin

-- Add new columns
ALTER TABLE users
    ADD COLUMN IF NOT EXISTS name         VARCHAR(64)     NOT NULL,
    ADD COLUMN IF NOT EXISTS surname      VARCHAR(64)     NOT NULL,
    ADD COLUMN IF NOT EXISTS gender       VARCHAR(6)      NOT NULL,
    ADD COLUMN IF NOT EXISTS birthday     TIMESTAMP   NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Drop the added columns
ALTER TABLE users
    DROP COLUMN IF EXISTS COLUMN name,
    DROP COLUMN IF EXISTS COLUMN surname,
    DROP COLUMN IF EXISTS COLUMN gender,
    DROP COLUMN IF EXISTS COLUMN birthday;
-- +goose StatementEnd