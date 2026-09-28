-- name: get_clients
SELECT * FROM clients;

-- name: get_clients_by_id
SELECT * FROM clients where id = ?;

-- name: insert_client
INSERT INTO clients (name, is_paying_me)
VALUES (?, ?)
RETURNING id;

-- name: delete_client
DELETE FROM clients WHERE id = ?
RETURNING id;