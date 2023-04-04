-- name: GetSensorById :one
SELECT feed_key, kind FROM sensors WHERE id = ?;

-- name: GetSensorsByType :many
SELECT id, feed_key FROM sensors WHERE kind = ?;