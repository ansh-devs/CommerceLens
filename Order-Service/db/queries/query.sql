-- name: CreateOrder :one
INSERT INTO orders (
  id, product_id, user_id, total_cost, status, fullname, address, product_name, description, price, created_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: ChangeOrderStatusById :exec
UPDATE orders SET status = $2
WHERE id = $1;

-- name: GetOrderById :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetAllOrdersByUserId :many
SELECT * FROM orders
WHERE id = $1 LIMIT 1;;
