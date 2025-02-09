CREATE TABLE users
(
    id            serial        PRIMARY KEY,
    username      VARCHAR(50)   NOT NULL UNIQUE,
    password      VARCHAR(100)  NOT NULL,
    email         VARCHAR(50)   NOT NULL UNIQUE,
    created_at    TIMESTAMP     DEFAULT (now() at time zone 'Europe/Athens')
);

CREATE TABLE deposits
(
    id            serial        PRIMARY KEY,
    issuer        VARCHAR(50)       REFERENCES users(username) ON DELETE CASCADE,
    amount        INTEGER       NOT NULL,
    description   VARCHAR(100)  NOT NULL,
    created_at    TIMESTAMP     DEFAULT (now() at time zone 'Europe/Athens')
);

CREATE TABLE deposit_members
(
    deposit_id    INTEGER       REFERENCES deposits(id) ON DELETE CASCADE,
    user_id       INTEGER       REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY (deposit_id, user_id)
);

CREATE TABLE withdraws
(
    id            serial        PRIMARY KEY,
    issuer        INTEGER       REFERENCES users(id) ON DELETE CASCADE,
    deposit_id       INTEGER       REFERENCES deposits(id) ON DELETE CASCADE,
    amount        INTEGER       NOT NULL,
    description   VARCHAR(100)  NOT NULL,
    created_at    TIMESTAMP     DEFAULT (now() at time zone 'Europe/Athens')
);
