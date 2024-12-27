CREATE TABLE IF NOT EXISTS ideas(
    id INTEGER PRIMARY KEY,
    image TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    link TEXT,
    tags TEXT
);
