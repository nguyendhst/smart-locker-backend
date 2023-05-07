// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: locker_user.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createLockerUser = `-- name: CreateLockerUser :execresult
INSERT INTO locker_user (user_id, locker_id) VALUES (?, ?)
`

type CreateLockerUserParams struct {
	UserID   int32 `json:"userID"`
	LockerID int32 `json:"lockerID"`
}

func (q *Queries) CreateLockerUser(ctx context.Context, arg CreateLockerUserParams) (sql.Result, error) {
	return q.exec(ctx, q.createLockerUserStmt, createLockerUser, arg.UserID, arg.LockerID)
}

const deleteLockerUser = `-- name: DeleteLockerUser :exec
DELETE FROM locker_user WHERE id = ?
`

func (q *Queries) DeleteLockerUser(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deleteLockerUserStmt, deleteLockerUser, id)
	return err
}

const deleteLockerUserFromUserIDAndLockerID = `-- name: DeleteLockerUserFromUserIDAndLockerID :execresult
DELETE FROM locker_user WHERE user_id = ? AND locker_id = ?
`

type DeleteLockerUserFromUserIDAndLockerIDParams struct {
	UserID   int32 `json:"userID"`
	LockerID int32 `json:"lockerID"`
}

func (q *Queries) DeleteLockerUserFromUserIDAndLockerID(ctx context.Context, arg DeleteLockerUserFromUserIDAndLockerIDParams) (sql.Result, error) {
	return q.exec(ctx, q.deleteLockerUserFromUserIDAndLockerIDStmt, deleteLockerUserFromUserIDAndLockerID, arg.UserID, arg.LockerID)
}

const getEntryFromUserIDAndLockerID = `-- name: GetEntryFromUserIDAndLockerID :one
SELECT id, user_id, locker_id FROM locker_user WHERE user_id = ? AND locker_id = ?
`

type GetEntryFromUserIDAndLockerIDParams struct {
	UserID   int32 `json:"userID"`
	LockerID int32 `json:"lockerID"`
}

type GetEntryFromUserIDAndLockerIDRow struct {
	ID       int32 `json:"id"`
	UserID   int32 `json:"userID"`
	LockerID int32 `json:"lockerID"`
}

func (q *Queries) GetEntryFromUserIDAndLockerID(ctx context.Context, arg GetEntryFromUserIDAndLockerIDParams) (GetEntryFromUserIDAndLockerIDRow, error) {
	row := q.queryRow(ctx, q.getEntryFromUserIDAndLockerIDStmt, getEntryFromUserIDAndLockerID, arg.UserID, arg.LockerID)
	var i GetEntryFromUserIDAndLockerIDRow
	err := row.Scan(&i.ID, &i.UserID, &i.LockerID)
	return i, err
}

const getLockersOfUser = `-- name: GetLockersOfUser :many
SELECT locker_id FROM locker_user WHERE user_id = ?
`

func (q *Queries) GetLockersOfUser(ctx context.Context, userID int32) ([]int32, error) {
	rows, err := q.query(ctx, q.getLockersOfUserStmt, getLockersOfUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var locker_id int32
		if err := rows.Scan(&locker_id); err != nil {
			return nil, err
		}
		items = append(items, locker_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateLockerUser = `-- name: UpdateLockerUser :execresult
UPDATE locker_user SET user_id = ?, locker_id = ? WHERE id = ?
`

type UpdateLockerUserParams struct {
	UserID   int32 `json:"userID"`
	LockerID int32 `json:"lockerID"`
	ID       int32 `json:"id"`
}

func (q *Queries) UpdateLockerUser(ctx context.Context, arg UpdateLockerUserParams) (sql.Result, error) {
	return q.exec(ctx, q.updateLockerUserStmt, updateLockerUser, arg.UserID, arg.LockerID, arg.ID)
}
