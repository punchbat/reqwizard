-- +goose Up
-- +goose StatementBegin
ALTER TABLE applications
    ADD IF NOT EXISTS file_path VARCHAR;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE applications
    DROP file_path;
-- +goose StatementEnd