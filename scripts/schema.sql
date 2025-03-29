CREATE TABLE users
(
    id            serial        PRIMARY KEY,
    username      VARCHAR(50)   NOT NULL UNIQUE,
    password      VARCHAR(100)  NOT NULL,
    email         VARCHAR(50)   NOT NULL UNIQUE,
    created_at    TIMESTAMP     DEFAULT (now() at time zone 'Europe/Athens')
);

CREATE TABLE sessions
(
    id              serial        PRIMARY KEY,
    issuer          INTEGER       REFERENCES users(id) ON DELETE CASCADE,
    amount          INTEGER       NOT NULL,
    withdraw_amount INTEGER       NOT NULL,
    description     VARCHAR(100)  NOT NULL,
    created_at      TIMESTAMP     DEFAULT (now() at time zone 'Europe/Athens')
);

CREATE TABLE session_members
(
    session_id    INTEGER       REFERENCES sessions(id) ON DELETE CASCADE,
    user_id       INTEGER       REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (session_id, user_id)
);

CREATE INDEX idx_sessions_issuer ON sessions (issuer);
CREATE INDEX idx_sessions_date ON sessions (created_at);
CREATE INDEX idx_sessions_created_at ON sessions(created_at);
CREATE INDEX idx_session_members_session_id ON session_members (session_id);
CREATE INDEX idx_session_members_user_id ON session_members(user_id);
