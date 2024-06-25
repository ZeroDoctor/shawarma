CREATE TABLE IF NOT EXISTS users (
    "name"         TEXT,
    "session"      TEXT UNIQUE,
    created_at     INT,
    modified_at    INT,

    PRIMARY KEY("name")
);