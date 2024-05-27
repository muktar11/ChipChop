-- name: CreateEstimation :exec
INSERT INTO estimations (product_name, price, location, user_id)
VALUES ($1, $2, $3, $4);

-- name: GetEstimationByID :one
SELECT id, product_name, price, location, user_id
FROM estimations
WHERE id = $1;

-- name: ListEstimations :many
SELECT id, product_name, price, location, user_id
FROM estimations
LIMIT $1 OFFSET $2;

-- name: UpdateEstimation :exec
UPDATE estimations
SET price = COALESCE($2, price)
WHERE id = $1;

-- name: DeleteEstimation :exec
DELETE FROM estimations
WHERE id = $1;
