-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :execresult
INSERT INTO users (name, password_hashed, email) VALUES ($1, $2, $3);

-- name: UpdateUser :execresult
UPDATE users SET name = $1, password_hashed = $2, email = $3 WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;