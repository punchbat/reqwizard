-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD IF NOT EXISTS application_created_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP application_created_at;
-- +goose StatementEnd
