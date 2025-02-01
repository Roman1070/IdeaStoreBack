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
