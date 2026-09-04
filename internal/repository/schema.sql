CREATE TABLE IF NOT EXISTS clients (
	id INTEGER PRIMARY KEY,
name TEXT NOT NULL UNIQUE,
is_paying_me BOOLEAN NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS projects (
	id INTEGER PRIMARY KEY,
name TEXT NOT NULL UNIQUE,
alias TEXT UNIQUE,
color TEXT,
code TEXT UNIQUE,
description TEXT,
client_id INTEGER,
FOREIGN KEY (client_id) REFERENCES clients(id) ON
DELETE
	RESTRICT
);

CREATE TABLE IF NOT EXISTS time_blocks (
	id INTEGER PRIMARY KEY,
project_id INTEGER NOT NULL,
started_at DATETIME NOT NULL,
stopped_at DATETIME,
notes TEXT,
FOREIGN KEY (project_id) REFERENCES projects(id) ON
DELETE
	RESTRICT,
	CHECK (stopped_at IS NULL
		OR stopped_at >= started_at)
);

