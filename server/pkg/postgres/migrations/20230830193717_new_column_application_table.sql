-- +goose Up
-- +goose StatementBegin
ALTER TABLE applications
    ADD COLUMN IF NOT EXISTS ticket_response_id UUID;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE applications
    DROP COLUMN IF EXISTS ticket_response_id;
-- +goose StatementEnd
