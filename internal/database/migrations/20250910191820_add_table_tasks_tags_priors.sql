-- +goose Up
CREATE TABLE priorities (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  color TEXT
);

CREATE TABLE tags (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  color TEXT
);

CREATE TABLE tasks (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  description TEXT,
  status INTEGER,
  priority_id INTEGER REFERENCES priorities(id),
  start_date DATETIME,
  due_date DATETIME,
  created_date DATETIME
);

CREATE TABLE task_tags (
  task_id INTEGER NOT NULL REFERENCES tasks(id),
  tag_id  INTEGER NOT NULL REFERENCES tags(id),
  PRIMARY KEY(task_id, tag_id)
);

-- +goose Down
DROP TABLE IF EXISTS task_tags;
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS priorities;