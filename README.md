# DEPOSINATOR

This app was made in order to track multi-member deposits and withdrawals
when using the same casino account

## Run locally

### Create a .env file in project root

```bash
ENVIRONMENT: dev
HTTP_PORT: 8080
POSTGRES_HOST: 127.0.0.1:5432
POSTGRES_USER: foo
POSTGRES_PASS: bar
STORE_KEY: secret
```

```bash
make dev
```

## Clean postgres data

```bash
make clean
```

### Open psql session

```bash
# switch to postgres user
su - postgres
# connect to db
psql
# show tables
\dt+
```
