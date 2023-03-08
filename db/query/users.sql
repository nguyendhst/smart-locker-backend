-- name: GetUser :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (name, password_hashed, email) VALUES (?, ?, ?);

-- name: UpdateUser :execresult
UPDATE users SET name = ?, password_hashed = ?, email = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;