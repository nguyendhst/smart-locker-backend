-- name: GetUserByEmail :one
SELECT password_hashed FROM users WHERE email = ?;

-- name: CreateUser :execresult
INSERT INTO users (name, password_hashed, email) VALUES (?, ?, ?);

-- name: UpdateUser :execresult
UPDATE users SET name = ?, password_hashed = ?, email = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;