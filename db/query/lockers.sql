-- name: GetLocker :one
SELECT locker_number, location, last_accessed FROM lockers WHERE id = ?;

-- name: GetLockerByNfcSig :one
SELECT id, lock_status FROM lockers WHERE nfc_sig = ?;

-- name: GetLockerByLockerNumber :one
SELECT * FROM lockers WHERE locker_number = ?;

-- name: GetLockerByLockerNumberAndLocation :one
SELECT * FROM lockers WHERE locker_number = ? AND location = ?;

-- name: CreateLocker :execresult
INSERT INTO lockers (locker_number, location, status, nfc_sig) VALUES (?, ?, ?, ?);

-- name: UpdateLocker :execresult
UPDATE lockers SET locker_number = ?, location = ?, status = ?, nfc_sig = ? WHERE id = ?;

-- name: UpdateLockerStatus :execresult
UPDATE lockers SET status = ? WHERE id = ?;

-- name: UpdateLockerNfcSig :execresult
UPDATE lockers SET nfc_sig = ? WHERE id = ?;

-- name: DeleteLocker :exec
DELETE FROM lockers WHERE id = ?;

-- name: UpdateLockStatus :execresult
UPDATE lockers SET lock_status = ? WHERE id = ?;