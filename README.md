## Required

`docker, sqlc, golang-migrate(https://github.com/golang-migrate/migrate)`

## Commands

```
docker compose up
```

**OR**

```
docker build -t simplebank:latest .
docker run --rm --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest
```
