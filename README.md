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

## To setup aws oidc for github actions

- https://www.youtube.com/watch?v=mel6N62WZb0
- https://github.com/aws-actions/configure-aws-credentials#assuming-a-role
- https://github.com/marketplace/actions/amazon-ecr-login-action-for-github-actions#aws-credentials
