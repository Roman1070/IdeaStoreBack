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