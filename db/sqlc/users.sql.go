// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO users (name, password_hashed, email) VALUES ($1, $2, $3)
`

func (q *Queries) CreateUser(ctx context.Context) (sql.Result, error) {
	return q.exec(ctx, q.createUserStmt, createUser)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, password_hashed, email, role, created_at, last_modified FROM users WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PasswordHashed,
		&i.Email,
		&i.Role,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :execresult
UPDATE users SET name = $1, password_hashed = $2, email = $3 WHERE id = $4
`

func (q *Queries) UpdateUser(ctx context.Context) (sql.Result, error) {
	return q.exec(ctx, q.updateUserStmt, updateUser)
}