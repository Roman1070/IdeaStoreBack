CREATE TABLE IF NOT EXISTS ideas(
    id SERIAL PRIMARY KEY,
    image TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    link TEXT,
    tags TEXT,
    user_id INTEGER NOT NULL,
    likes_count INTEGER NOT NULL DEFAULT 0
);
