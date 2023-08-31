-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ticket_responses
(
    id              UUID                    NOT NULL PRIMARY KEY,
    application_id  UUID                    NOT NULL REFERENCES applications(id) ON DELETE CASCADE,
    user_id         UUID                    NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    manager_id      UUID                    NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    text            TEXT                    NOT NULL,
    created_at      TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,
    deleted_at      TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS ticket_responses;
-- +goose StatementEnd