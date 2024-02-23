CREATE TABLE IF NOT EXISTS runners (
    uuid        TEXT UNIQUE, -- I know, I know
    "type"      TEXT,
    hostname    TEXT,
    created_at  INT,
    modified_at INT,

    PRIMARY KEY(hostname, "type")
);

CREATE TABLE IF NOT EXISTS logs (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    "level"    INT,
    "data"     TEXT,
    created_at INT,

    runner_id TEXT,

    FOREIGN KEY(runner_id) REFERENCES runners(uuid)
);