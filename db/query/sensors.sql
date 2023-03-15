-- name: GetSensorById :one
SELECT feed_key, type FROM sensors WHERE id = ?;