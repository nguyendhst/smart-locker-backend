// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: lockers.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createLocker = `-- name: CreateLocker :execresult
INSERT INTO lockers (locker_number, location, status, nfc_sig) VALUES ($1, $2, $3, $4)
`

func (q *Queries) CreateLocker(ctx context.Context) (sql.Result, error) {
	return q.exec(ctx, q.createLockerStmt, createLocker)
}

const deleteLocker = `-- name: DeleteLocker :exec
DELETE FROM lockers WHERE id = $1
`

func (q *Queries) DeleteLocker(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteLockerStmt, deleteLocker)
	return err
}

const getLocker = `-- name: GetLocker :one
SELECT id, locker_number, location, status, nfc_sig, created_at, last_modified FROM lockers WHERE id = $1
`

func (q *Queries) GetLocker(ctx context.Context) (Locker, error) {
	row := q.queryRow(ctx, q.getLockerStmt, getLocker)
	var i Locker
	err := row.Scan(
		&i.ID,
		&i.LockerNumber,
		&i.Location,
		&i.Status,
		&i.NfcSig,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}

const getLockerByLockerNumber = `-- name: GetLockerByLockerNumber :one
SELECT id, locker_number, location, status, nfc_sig, created_at, last_modified FROM lockers WHERE locker_number = $1
`

func (q *Queries) GetLockerByLockerNumber(ctx context.Context) (Locker, error) {
	row := q.queryRow(ctx, q.getLockerByLockerNumberStmt, getLockerByLockerNumber)
	var i Locker
	err := row.Scan(
		&i.ID,
		&i.LockerNumber,
		&i.Location,
		&i.Status,
		&i.NfcSig,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}

const getLockerByLockerNumberAndLocation = `-- name: GetLockerByLockerNumberAndLocation :one
SELECT id, locker_number, location, status, nfc_sig, created_at, last_modified FROM lockers WHERE locker_number = $1 AND location = $2
`

func (q *Queries) GetLockerByLockerNumberAndLocation(ctx context.Context) (Locker, error) {
	row := q.queryRow(ctx, q.getLockerByLockerNumberAndLocationStmt, getLockerByLockerNumberAndLocation)
	var i Locker
	err := row.Scan(
		&i.ID,
		&i.LockerNumber,
		&i.Location,
		&i.Status,
		&i.NfcSig,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}

const getLockerByNfcSig = `-- name: GetLockerByNfcSig :one
SELECT id, locker_number, location, status, nfc_sig, created_at, last_modified FROM lockers WHERE nfc_sig = $1
`

func (q *Queries) GetLockerByNfcSig(ctx context.Context) (Locker, error) {
	row := q.queryRow(ctx, q.getLockerByNfcSigStmt, getLockerByNfcSig)
	var i Locker
	err := row.Scan(
		&i.ID,
		&i.LockerNumber,
		&i.Location,
		&i.Status,
		&i.NfcSig,
		&i.CreatedAt,
		&i.LastModified,
	)
	return i, err
}

const updateLocker = `-- name: UpdateLocker :execresult
UPDATE lockers SET locker_number = $1, location = $2, status = $3, nfc_sig = $4 WHERE id = $5
`

func (q *Queries) UpdateLocker(ctx context.Context) (sql.Result, error) {
	return q.exec(ctx, q.updateLockerStmt, updateLocker)
}

const updateLockerNfcSig = `-- name: UpdateLockerNfcSig :execresult
UPDATE lockers SET nfc_sig = $1 WHERE id = $2
`

func (q *Queries) UpdateLockerNfcSig(ctx context.Context) (sql.Result, error) {
	return q.exec(ctx, q.updateLockerNfcSigStmt, updateLockerNfcSig)
}

const updateLockerStatus = `-- name: UpdateLockerStatus :execresult
UPDATE lockers SET status = $1 WHERE id = $2
`

func (q *Queries) UpdateLockerStatus(ctx context.Context) (sql.Result, error) {
	return q.exec(ctx, q.updateLockerStatusStmt, updateLockerStatus)
}
