-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS roles
(
    id                  UUID                NOT NULL PRIMARY KEY,
    name                CITEXT              NOT NULL UNIQUE,
    created_at          TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP,
    deleted_at          TIMESTAMP
);

INSERT INTO roles (id, name) VALUES
    ('d76f6b25-e67a-4d5a-bb17-7266b7a6f8d0', 'user'),
    ('f6978b84-89f6-4b88-8d44-07fc79a3e9a5', 'manager');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
