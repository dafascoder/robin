-- name: CreateAccount :one
INSERT INTO account (email, password) VALUES ($1,$2) RETURNING id,"email";

-- name: GetAccountByEmail :one
SELECT * FROM account
WHERE email = $1 LIMIT 1;

-- name: VerifyAccount :one
UPDATE account SET verified = true WHERE id = $1
RETURNING id, email, verified;

-- name: UpdateRefreshTokenVersion :one
UPDATE account SET refresh_token_version = refresh_token_version + 1 WHERE id = $1
RETURNING id, email, refresh_token_version;

-- name: UpdatePassword :one
UPDATE account SET password = $2 WHERE id = $1
RETURNING id, email;
