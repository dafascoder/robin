-- name: CreateAccount :one
INSERT INTO account (email, password) VALUES ($1,$2) RETURNING id,"email";

-- name: GetAccountByEmail :one
SELECT * FROM account
WHERE email = $1 LIMIT 1;
