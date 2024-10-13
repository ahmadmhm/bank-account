-- Account model
-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY owner;

-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateAccount :exec
UPDATE accounts
  set owner = $2, balance = $3, currency = $4
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- Entry model
-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY account_id;

-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: UpdateEntry :exec
UPDATE entries
  set account_id = $2, amount = $3
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

-- Transfer model
-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY from_account_id;

-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: UpdateTransfer :exec
UPDATE transfers
  set from_account_id = $2, to_account_id = $3, amount = $4
WHERE id = $1;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;