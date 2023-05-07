-- name: GetSensorsOfLocker :many
SELECT sensor_id FROM locker_sensor WHERE locker_id = ?;

-- name: CreateSensorLocker :execresult
INSERT INTO locker_sensor (locker_id, sensor_id) VALUES (?, ?);