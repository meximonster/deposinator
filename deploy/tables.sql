CREATE TABLE accounts
(
    id            serial        PRIMARY KEY,
    username      VARCHAR(20)   NOT NULL,
    password      VARCHAR(20)   NOT NULL,
    email         VARCHAR(10)   NOT NULL,
    created_at    TIMESTAMP     DEFAULT (now() at time zone 'Europe/Athens')
);
