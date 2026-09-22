-- name: get_projects
SELECT * FROM projects;

-- name: get_projects_by_id
SELECT * FROM projects WHERE id = ?;

-- name: insert_project
INSERT INTO projects (name, alias, color, code, description, client_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: delete_project
DELETE FROM projects WHERE id = ?
RETURNING id;
