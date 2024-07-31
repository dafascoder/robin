-- +goose Up
ALTER TABLE "user"
    ADD FOREIGN KEY (account_id) REFERENCES account (id) ON DELETE CASCADE;
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
ALTER TABLE "user" DROP CONSTRAINT "user_account_id_fkey";
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
