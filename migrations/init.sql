CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    pass_hash TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_email ON users(email);


CREATE TABLE IF NOT EXISTS boards(
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    ideas_ids TEXT NOT NULL,
    user_id INTEGER NOT NULL
);



CREATE TABLE IF NOT EXISTS chats(
    id SERIAL PRIMARY KEY,
    first_id INTEGER NOT NULL,
    second_id INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS  ids_index ON chats(first_id,second_id);

CREATE TABLE IF NOT EXISTS messages(
    id SERIAL PRIMARY KEY,
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



ALTER TABLE messages
ADD COLUMN idea_id INTEGER NOT NULL DEFAULT -1;



CREATE TABLE IF NOT EXISTS comments(
    id SERIAL PRIMARY KEY NOT NULL,
    idea_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    creation_date TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_idea ON comments(idea_id);



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



CREATE TABLE IF NOT EXISTS profiles(
    id SERIAL PRIMARY KEY NOT NULL,
    email TEXT,
    avatarImage TEXT,
    name TEXT,
    description TEXT,
    link TEXT,
    boards TEXT,
    savedIdeas TEXT
);

CREATE INDEX IF NOT EXISTS names_idx ON profiles(name);
CREATE INDEX IF NOT EXISTS emails_idx ON profiles(email);



ALTER TABLE profiles
    ADD COLUMN liked_ideas TEXT NOT NULL;