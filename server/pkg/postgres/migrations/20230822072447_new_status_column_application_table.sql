-- +goose Up
-- +goose StatementBegin
CREATE TYPE application_status AS ENUM ('canceled', 'waiting', 'working', 'done');

ALTER TABLE applications
    ADD COLUMN IF NOT EXISTS status application_status NOT NULL DEFAULT 'waiting';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE applications
    DROP COLUMN IF EXISTS status;
-- +goose StatementEnd