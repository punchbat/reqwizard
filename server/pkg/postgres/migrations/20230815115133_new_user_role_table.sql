-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_roles
(
    id          UUID                NOT NULL PRIMARY KEY,
    user_id     UUID                NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role_id     UUID                NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    status      VARCHAR             NOT NULL,
    created_at  TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd
