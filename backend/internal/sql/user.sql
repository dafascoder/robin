-- name: CreateUser :one
INSERT INTO "user" (account_id, name) VALUES ($1,$2) RETURNING id, account_id, name;

-- name: GetUserByAccountID :one
SELECT * FROM "user"
WHERE account_id = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM "user"
WHERE id = $1 LIMIT 1;