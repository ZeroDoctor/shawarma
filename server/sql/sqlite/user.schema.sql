CREATE TABLE IF NOT EXISTS users (
    "name"         TEXT,
    "session"      TEXT UNIQUE,
    git_remote     TEXT,
    created_at     INT,
    modified_at    INT,

    PRIMARY KEY("name")
);