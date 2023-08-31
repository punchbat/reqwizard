-- +goose Up
-- +goose StatementBegin
ALTER TABLE applications
    DROP COLUMN IF EXISTS file_path;

ALTER TABLE applications
    ADD COLUMN IF NOT EXISTS file_name VARCHAR;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE applications
    ADD COLUMN IF NOT EXISTS file_path VARCHAR;

ALTER TABLE applications
    DROP COLUMN IF EXISTS file_name;
-- +goose StatementEnd
