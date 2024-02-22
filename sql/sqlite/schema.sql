CREATE TABLE IF NOT EXISTS organizations (
    id          INT,
    "owner"     TEXT,
    "name"      TEXT,
    created_at  INT,
    modified_at INT,

    PRIMARY KEY("owner", "name")
);

CREATE TABLE IF NOT EXISTS repositories (
    id          INT,
    "owner"     TEXT,
    "name"      TEXT,
    created_at  INT,
    modified_at INT,

    PRIMARY KEY("owner", "name")
);

CREATE TABLE IF NOT EXISTS branches (
    id               INTEGER PRIMARY KEY AUTOINCREMENT,
    "name"           TEXT,
    latest_commit    TEXT,
    created_at       INT,
    modified_at      INT,

    repo_id          INT,

    FOREIGN KEY(repo_id) REFERENCES repositories(id),
    FOREIGN KEY(latest_commit) REFERENCES commits("commit")
);

CREATE TABLE IF NOT EXISTS commits (
    "hash"   TEXT PRIMARY KEY,
    author     TEXT,
    created_at INT,

    branch_id  INT,

    FOREIGN KEY(branch_id) REFERENCES branches(id)
);

CREATE TABLE IF NOT EXISTS runners (
    "type"      TEXT,
    hostname    TEXT,
    created_at  INT,
    modified_at INT,

    pipeline_id INT,

    FOREIGN KEY (pipeline_id) REFERENCES pipelines(id),
    PRIMARY KEY(hostname, "type")
);

CREATE TABLE IF NOT EXISTS pipelines (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    "type"      TEXT,
    "status"    TEXT,
    created_at  INT,
    modified_at INT,

    repo_id     INT,

    FOREIGN KEY(repo_id) REFERENCES repositories(id)
);

CREATE TABLE IF NOT EXISTS steps (
    id          INT,
    "name"      TEXT,
    "image"     TEXT,
    commands    TEXT,
    privileged  INT2,
    detach      INT2,
    created_at  INT,
    modified_at INT,

    pipeline_id INT,

    FOREIGN KEY (pipeline_id) REFERENCES pipelines(id),
    PRIMARY KEY("name", pipeline_id)
);

CREATE TABLE IF NOT EXISTS environments (
    "key"       TEXT,
    "data"      TEXT,
    protected   INT2,
    created_at  INT,
    modified_at INT,

    repo_id     INT,
    org_id      INT,

    FOREIGN KEY (repo_id) REFERENCES repositories(id),
    FOREIGN KEY (org_id) REFERENCES organizations(id),
    PRIMARY KEY("key", repo_id, org_id)
);

CREATE TABLE IF NOT EXISTS events (
    webhook     TEXT,
    "type"      TEXT,
    "action"    TEXT,
    deadline    TEXT,
    created_at  INT,
    modified_at INT,

    pipeline_id INT,
    step_id     INT,

    FOREIGN KEY (pipeline_id) REFERENCES pipelines(id),
    FOREIGN KEY (step_id) REFERENCES steps(id),
    PRIMARY KEY ("type", pipeline_id, step_id)
);

CREATE TABLE IF NOT EXISTS logs (
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    "level"    INT,
    "data"     TEXT,
    created_at INT
);
