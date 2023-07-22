## Required

`docker, sqlc, golang-migrate(https://github.com/golang-migrate/migrate)`

## Commands

```
docker build -t simplebank:latest .
docker run --rm --name simplebank -p 8080:8080 simplebank:latest
docker run --rm --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest
```
