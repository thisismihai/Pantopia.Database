-- name: CreateAccount :one
INSERT INTO accounts (
    owner,
    company_name,
    industry,
    target_industry,
    target_position,
    target_size
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;


-- name: GetAccount :one
Select * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET company_name = $2
where id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;