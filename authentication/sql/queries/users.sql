-- name: InsertUser :one
INSERT INTO users 
        (email, first_name, last_name, password, active, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7) 
RETURNING *;

-- name: GetAll :many
SELECT * FROM users 
ORDER BY last_name;

-- name: GetByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :one
UPDATE users SET
email = $1,
first_name = $2,
last_name = $3,
active = $4,
updated_at = $5
WHERE id = $6
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users 
SET password = $1 
WHERE id = $2;