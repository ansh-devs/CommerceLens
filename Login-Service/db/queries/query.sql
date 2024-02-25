-- name: CreateUser :one
INSERT INTO users (
  id, fullname,email,password,address,created_at
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CheckUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;


-- name: CheckUserIsRegistered :one
SELECT count(*) FROM users
WHERE email = $1 LIMIT 1;