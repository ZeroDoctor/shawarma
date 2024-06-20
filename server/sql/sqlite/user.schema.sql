CREATE TABLE IF NOT EXISTS users (
    "name"         TEXT,
    "session"      TEXT UNIQUE,
    github_token   TEXT,
    github_user_id INT,
    created_at     INT,
    modified_at    INT,

    PRIMARY KEY("name")
    FOREIGN KEY(github_user_id) REFERENCES github_users(id),
);