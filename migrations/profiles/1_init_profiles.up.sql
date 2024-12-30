CREATE TABLE IF NOT EXISTS profiles(
    id INTEGER PRIMARY KEY NOT NULL,
    email TEXT NOT NULL,
    avatarImage TEXT,
    name TEXT NOT NULL,
    description TEXT,
    link TEXT,
    boards TEXT NOT NULL DEFAULT "",
    savedIdeas TEXT NOT NULL DEFAULT ""
);