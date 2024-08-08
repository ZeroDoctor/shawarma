CREATE TABLE IF NOT EXISTS users (
    uuid           TEXT UNIQUE,
    "name"         TEXT,
    "session"      TEXT UNIQUE,
    avatar_url     TEXT,
    git_remote     TEXT,
    is_owner       INT,
    created_at     INT,
    modified_at    INT,

    PRIMARY KEY("name")
);

CREATE TABLE IF NOT EXISTS organizations (
    uuid        TEXT UNIQUE, -- yes I know
    "owner"     TEXT,
    "name"      TEXT,
    avatar_url  TEXT,
    created_at  INT,
    modified_at INT,

    FOREIGN KEY("owner") REFERENCES users("name"),
    PRIMARY KEY("name")
);
