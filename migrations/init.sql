CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    pass_hash BLOB NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_email ON users(email);

CREATE TABLE IF NOT EXISTS apps(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    secret TEXT NOT NULL UNIQUE
);
ALTER TABLE users
    ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT FALSE;

CREATE TABLE IF NOT EXISTS boards(
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    ideas_ids TEXT NOT NULL,
    user_id INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS chats(
    id INTEGER PRIMARY KEY,
    first_id INTEGER NOT NULL,
    second_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS  ids_index ON chats(first_id,second_id);

CREATE TABLE IF NOT EXISTS messages(
    id INTEGER PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    reciever_id INTEGER NOT NULL,
    file_name TEXT,
    content TEXT NOT NULL,
    send_date TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS sender_ids_index ON messages(sender_id);
CREATE INDEX IF NOT EXISTS  reciever_ids_index ON messages(reciever_id);
ALTER TABLE messages
ADD COLUMN sending_date_seconds INTEGER NOT NULL DEFAULT 0;


CREATE TABLE IF NOT EXISTS comments(
    id INTEGER PRIMARY KEY NOT NULL,
    idea_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    creation_date TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_idea ON comments(idea_id);


CREATE TABLE IF NOT EXISTS ideas(
    id INTEGER PRIMARY KEY,
    image TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    link TEXT,
    tags TEXT,
    user_id INTEGER NOT NULL,
    likes_count INTEGER NOT NULL DEFAULT 0
);


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