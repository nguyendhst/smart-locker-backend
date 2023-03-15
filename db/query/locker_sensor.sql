-- name: GetSensorsOfLocker :many
SELECT sensor_id FROM locker_sensor WHERE locker_id = ?;