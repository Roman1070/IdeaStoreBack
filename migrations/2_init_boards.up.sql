CREATE TABLE IF NOT EXISTS boards(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    ideas_ids TEXT NOT NULL,
    user_id INTEGER NOT NULL
);