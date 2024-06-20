CREATE TABLE IF NOT EXISTS github_users (
    id                INT PRIMARY KEY,
    avatar_url        TEXT,
    gravatar_id       TEXT,
    "url"             TEXT,
    organizations_url TEXT,
    repos_url         TEXT,
    "type"            TEXT,
    "name"            TEXT UNIQUE NOT NULL,
    created_at        TEXT,
    updated_at        TEXT
);

CREATE TABLE IF NOT EXISTS organizations (
    uuid        TEXT UNIQUE, -- yes I know
    "owner"     TEXT,
    "name"      TEXT,
    created_at  INT,
    modified_at INT,

    PRIMARY KEY("owner", "name")
);

CREATE TABLE IF NOT EXISTS repositories (
    uuid        TEXT UNIQUE, -- hear me out
    "owner"     TEXT,
    "name"      TEXT,
    created_at  INT,
    modified_at INT,

    org_id      TEXT,

    FOREIGN KEY(org_id) REFERENCES organizations(uuid),
    PRIMARY KEY("owner", "name")
);

CREATE TABLE IF NOT EXISTS branches (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    "name"           TEXT,
    created_at       INT,
    modified_at      INT,

    latest_commit    TEXT,
    repo_id          INT,

    FOREIGN KEY(repo_id) REFERENCES repositories(id),
    FOREIGN KEY(latest_commit) REFERENCES commits("commit")
);

CREATE TABLE IF NOT EXISTS commits (
    "hash"     TEXT PRIMARY KEY,
    author     TEXT,
    created_at INT,

    branch_id  INT,

    FOREIGN KEY(branch_id) REFERENCES branches(id)
);

