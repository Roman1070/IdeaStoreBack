CREATE TABLE IF NOT EXISTS boards(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    ideas_ids TEXT NOT NULL
);