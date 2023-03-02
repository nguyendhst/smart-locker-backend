-- name: GetLocker :one
SELECT * FROM lockers WHERE id = $1;

-- name: GetLockerByNfcSig :one
SELECT * FROM lockers WHERE nfc_sig = $1;

-- name: GetLockerByLockerNumber :one
SELECT * FROM lockers WHERE locker_number = $1;

-- name: GetLockerByLockerNumberAndLocation :one
SELECT * FROM lockers WHERE locker_number = $1 AND location = $2;

-- name: CreateLocker :execresult
INSERT INTO lockers (locker_number, location, status, nfc_sig) VALUES ($1, $2, $3, $4);

-- name: UpdateLocker :execresult
UPDATE lockers SET locker_number = $1, location = $2, status = $3, nfc_sig = $4 WHERE id = $5;

-- name: UpdateLockerStatus :execresult
UPDATE lockers SET status = $1 WHERE id = $2;

-- name: UpdateLockerNfcSig :execresult
UPDATE lockers SET nfc_sig = $1 WHERE id = $2;

-- name: DeleteLocker :exec
DELETE FROM lockers WHERE id = $1;
