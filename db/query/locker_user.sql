-- name: GetLockerUser :one
SELECT * FROM locker_user WHERE id = $1;

-- name: CreateLockerUser :execresult
INSERT INTO locker_user (user_id, locker_id) VALUES ($1, $2);

-- name: UpdateLockerUser :execresult
UPDATE locker_user SET user_id = $1, locker_id = $2 WHERE id = $3;

-- name: DeleteLockerUser :exec
DELETE FROM locker_user WHERE id = $1;
