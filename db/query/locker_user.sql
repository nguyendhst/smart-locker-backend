-- name: GetLockersOfUser :many
SELECT locker_id FROM locker_user WHERE user_id = ?;

-- name: CreateLockerUser :execresult
INSERT INTO locker_user (user_id, locker_id) VALUES (?, ?);

-- name: UpdateLockerUser :execresult
UPDATE locker_user SET user_id = ?, locker_id = ? WHERE id = ?;

-- name: DeleteLockerUser :exec
DELETE FROM locker_user WHERE id = ?;

-- name: GetEntryFromUserIDAndLockerID :one
SELECT id, user_id, locker_id FROM locker_user WHERE user_id = ? AND locker_id = ?;

-- name: DeleteLockerUserFromUserIDAndLockerID :execresult
DELETE FROM locker_user WHERE user_id = ? AND locker_id = ?;