CREATE TABLE IF NOT EXISTS users (
    "name"       TEXT,
    "session"    TEXT,
    github_token TEXT,
    created_at   INT,
    modified_at  INT,

    PRIMARY KEY("name", "session")
);