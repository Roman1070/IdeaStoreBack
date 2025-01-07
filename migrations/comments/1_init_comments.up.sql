CREATE TABLE IF NOT EXISTS comments(
    id INTEGER PRIMARY KEY NOT NULL,
    idea_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    creation_date TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_idea ON comments(idea_id);