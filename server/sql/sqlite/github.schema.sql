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

CREATE TABLE IF NOT EXISTS github_orgs (
    id                 INT PRIMARY KEY,
    "url"              TEXT,
    repos_url          TEXT,
    hooks_url          TEXT,
    issues_url         TEXT,
    members_url        TEXT,
    public_members_url TEXT,
    avatar_url         TEXT,
    "description"      TEXT,
    "name"             TEXT,
    company            TEXT,
    created_at         TEXT,
    updated_at         TEXT,
    archived_at        TEXT,
    "type"             TEXT,
);

CREATE TABLE IF NOT EXISTS github_users_orgs (
    github_user_id INT,
    github_org_id  INT,
    created_at     INT,

    PRIMARY KEY(github_user_id, github_org_id),
    FOREIGN KEY(github_user_id) REFERENCES github_users(id),
    FOREIGN KEY(github_org_id) REFERENCES github_orgs(id)
);