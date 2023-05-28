-- name: CreateLead :one
INSERT INTO leads (
    account_id,
    email,
    telephone_number,
    target_email,
    company_name,
    conversation
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetLead :one
SELECT * FROM leads
WHERE id = $1 LIMIT 1;

-- name: ListLeads :many
SELECT * FROM leads
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateLead :one
UPDATE leads
SET email = $2, 
    telephone_number = $3,
    target_email = $4,
    company_name = $5,
    conversation = $6
WHERE id = $1
RETURNING *;

-- name: DeleteLead :exec
DELETE FROM leads
WHERE id = $1;


