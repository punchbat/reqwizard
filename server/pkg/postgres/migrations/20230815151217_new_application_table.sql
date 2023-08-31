-- +goose Up
-- +goose StatementBegin
CREATE TYPE application_type AS ENUM ('financial', 'general');
CREATE TYPE application_sub_type AS ENUM ('information', 'account_help', 'refunds', 'payment');

CREATE TABLE IF NOT EXISTS applications
(
    id              UUID                    NOT NULL PRIMARY KEY,
    user_id         UUID                    NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    manager_id      UUID                    REFERENCES users(id) ON DELETE CASCADE,
    type            application_type        NOT NULL,
    sub_type        application_sub_type    NOT NULL,
    title           VARCHAR                 NOT NULL,
    description     VARCHAR                 NOT NULL,
    created_at      TIMESTAMP               NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP,
    deleted_at      TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS applications;
-- +goose StatementEnd
