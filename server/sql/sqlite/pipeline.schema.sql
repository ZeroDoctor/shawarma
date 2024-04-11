CREATE TABLE IF NOT EXISTS pipelines (
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    "type"      TEXT,
    "status"    TEXT,
    created_at  INT,
    modified_at INT,

    repo_id     TEXT,
    runner_id   TEXT,

    FOREIGN KEY(repo_id) REFERENCES repositories(uuid),
    FOREIGN KEY(runner_id) REFERENCES runners(uuid)
);

CREATE TABLE IF NOT EXISTS steps (
    uuid        TEXT UNIQUE,
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

    repo_id     TEXT,
    org_id      TEXT,

    FOREIGN KEY (repo_id) REFERENCES repositories(uuid),
    FOREIGN KEY (org_id) REFERENCES organizations(uuid),
    PRIMARY KEY("key", repo_id, org_id)
);

CREATE TABLE IF NOT EXISTS events (
    webhook     TEXT,
    "type"      TEXT,
    status_name TEXT,
    "action"    TEXT,
    deadline    TEXT,
    "after"     TEXT,
    created_at  INT,
    modified_at INT,

    pipeline_id INT,
    step_id     TEXT,

    FOREIGN KEY (pipeline_id) REFERENCES pipelines(id),
    FOREIGN KEY (step_id) REFERENCES steps(uuid),
    PRIMARY KEY ("type", pipeline_id, step_id)
);
