CREATE TABLE IF NOT EXISTS repositories (
    uuid           TEXT UNIQUE, -- hear me out
    "owner"        TEXT,
    "name"         TEXT,
    default_branch TEXT,
    active         INT,
    created_at     INT,
    modified_at    INT,
    owner_type     TEXT, -- either user || organization
    owner_id       TEXT,

    PRIMARY KEY("owner", "name")
);

CREATE TABLE IF NOT EXISTS branches (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    "name"           TEXT,
    created_at       INT,
    modified_at      INT,

    repo_id          TEXT,

    FOREIGN KEY(repo_id) REFERENCES repositories(uuid)
);

CREATE TABLE IF NOT EXISTS commits (
    "hash"     TEXT PRIMARY KEY,
    author     TEXT,
    "message"  TEXT,
    created_at INT,

    branch_id  INT,

    FOREIGN KEY(branch_id) REFERENCES branches(id)
);

