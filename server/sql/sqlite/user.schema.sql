CREATE TABLE IF NOT EXISTS users (
    uuid           TEXT UNIQUE,
    "name"         TEXT,
    "session"      TEXT UNIQUE,
    avatar_url     TEXT,
    git_remote     TEXT,
    created_at     INT,
    modified_at    INT,

    PRIMARY KEY("name")
);