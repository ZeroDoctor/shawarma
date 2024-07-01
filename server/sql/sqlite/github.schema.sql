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
    user_id    INT,
    org_id     INT,
    created_at INT,

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
    default_branch    TEXT,

    FOREIGN KEY(owner_id) REFERENCES github_owners(id)
);

CREATE TABLE IF NOT EXISTS github_branches (
    "name"       TEXT,
    "url"        TEXT,
    author_id    INT,
    committer_id INT,
    repo_id      INT,
    sha          TEXT,

    FOREIGN KEY(sha) REFERENCES github_commits(sha),
    FOREIGN KEY(repo_id) REFERENCES github_repos(id),
    PRIMARY KEY(repo_id, "name")
);

CREATE TABLE IF NOT EXISTS github_commits (
    sha       TEXT PRIMARY KEY,
    "message" TEXT,
    "url"     TEXT
);

CREATE TABLE IF NOT EXISTS github_commit_parents (
    parent_sha TEXT,
    child_sha  TEXT,

    UNIQUE(parent_sha, child_sha),
    FOREIGN KEY(parent_sha) REFERENCES github_commits(sha),
    FOREIGN KEY(child_sha) REFERENCES github_commits(sha)
)
