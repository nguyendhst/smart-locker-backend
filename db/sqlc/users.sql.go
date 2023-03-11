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
INSERT INTO users (name, password_hashed, email) VALUES (?, ?, ?)
`

type CreateUserParams struct {
	Name           string `json:"name"`
	PasswordHashed string `json:"passwordHashed"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.createUserStmt, createUser, arg.Name, arg.PasswordHashed, arg.Email)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT password_hashed FROM users WHERE email = ?
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (string, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var password_hashed string
	err := row.Scan(&password_hashed)
	return password_hashed, err
}

const updateUser = `-- name: UpdateUser :execresult
UPDATE users SET name = ?, password_hashed = ?, email = ? WHERE id = ?
`

type UpdateUserParams struct {
	Name           string `json:"name"`
	PasswordHashed string `json:"passwordHashed"`
	Email          string `json:"email"`
	ID             int32  `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Name,
		arg.PasswordHashed,
		arg.Email,
		arg.ID,
	)
}
