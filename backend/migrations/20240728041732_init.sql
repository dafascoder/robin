-- +goose Up
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

    CREATE TABLE account (
        id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
        email TEXT NOT NULL UNIQUE,
        password text NOT NULL,
        created_at timestamp DEFAULT now(),
        updated_at timestamp  DEFAULT now()
    );
    CREATE TABLE "user" (
        id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
        name text NOT NULL,
        bio text,
        image text,
        created_at timestamp DEFAULT now(),
        updated_at timestamp DEFAULT now()
    );
    CREATE TABLE "organization" (
        id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
        name text NOT NULL,
        info text,
        image text,
        created_at timestamp DEFAULT now(),
        updated_at timestamp DEFAULT now()
    );
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
    DROP TABLE account;
    DROP TABLE "user";
    DROP TABLE "organization";
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd