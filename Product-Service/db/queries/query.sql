-- name: CreateProduct :one
INSERT INTO products (
  id, product_name,description,price,created_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetAllProducts :many
SELECT * FROM products;

-- name: CheckUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

---- name: CheckProductIsAdded :one
--SELECT count(*) FROM products
--WHERE email = $1 LIMIT 1;