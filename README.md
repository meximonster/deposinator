# DEPOSINATOR

Create a .env file in project root using docs/example.env

## Run locally

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
