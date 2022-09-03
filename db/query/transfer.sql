-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    status,
    amount
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers 
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET status = $2,
    amount = $3
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;