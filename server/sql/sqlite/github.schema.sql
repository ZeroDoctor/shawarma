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
    "type"             TEXT
);

CREATE TABLE IF NOT EXISTS github_owners (
    id                INT PRIMARY KEY,
    avatar_url        TEXT,
    gravatar_id       TEXT,
    "url"             TEXT,
    organizations_url TEXT,
    repos_url         TEXT,
    "type"            TEXT
);

CREATE TABLE IF NOT EXISTS github_users_orgs (
    user_id INT,
    org_id  INT,
    created_at     INT,

    PRIMARY KEY(user_id, org_id),
    FOREIGN KEY(user_id) REFERENCES github_users(id),
    FOREIGN KEY(org_id) REFERENCES github_orgs(id)
);

CREATE TABLE IF NOT EXISTS github_repos (
    id                INT PRIMARY KEY,
    owner_id          INT NOT NULL,
    "name"            TEXT,
    full_name         TEXT,
    "description"     TEXT,
    "url"             TEXT,
    collaborators_url TEXT,
    hooks_url         TEXT,
    issue_events_url  TEXT,
    branches_url      TEXT,
    tags_url          TEXT,
    statuses_url      TEXT,
    commits_url       TEXT,
    merges_url        TEXT,
    issues_url        TEXT,
    pulls_url         TEXT,
    created_at        TEXT,
    updated_at        TEXT,
    pushed_at         TEXT,
    has_issues        INT,
    archived          INT,
    open_issues_count INT,
    visibility        TEXT,

    FOREIGN KEY(owner_id) REFERENCES github_owners(id)
);
