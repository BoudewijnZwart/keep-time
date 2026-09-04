-- name: get_time_blocks
SELECT * FROM time_blocks;

-- name: get_time_blocks_by_id
SELECT * FROM time_blocks WHERE id = ?;

-- name: insert_time_block
INSERT INTO time_blocks (project_id, started_at, notes)
VALUES (?, ?, ?)
RETURNING id;

-- name: stop_time_block
UPDATE time_blocks
SET stopped_at = ?
WHERE id = ? AND stopped_at IS NULL
RETURNING id;

-- name: delete_time_block
DELETE FROM time_blocks WHERE id = ?
RETURNING id;
