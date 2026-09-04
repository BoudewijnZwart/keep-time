-- name: get_clients
SELECT * FROM clients;

-- name: get_clients_by_id
SELECT * FROM clients where id = ?;

-- name: get_projects
SELECT * FROM projects;

-- name: get_projects_by_id
SELECT * FROM projects WHERE id = ?;

-- name: get_time_blocks
SELECT * FROM time_blocks;

-- name: get_time_blocks_by_id
SELECT * FROM time_blocks WHERE id = ?;

-- name: insert_client
INSERT INTO clients (name, is_paying_me)
VALUES (?, ?)
RETURNING id;

-- name: insert_project
INSERT INTO projects (name, code, description, client_id)
VALUES (?, ?, ?, ?)
RETURNING id;

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

-- name: delete_project
DELETE FROM projects WHERE id = ?
RETURNING id;

-- name: delete_client
DELETE FROM clients WHERE id = ?
RETURNING id;
