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

## Secrets manager

```
aws secretsmanager get-secret-value --secret-id <friendly name/arn> --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]'
```

## ECR docker login

```
aws ecr get-login-password | docker login --username AWS --password-stdin aws_account_id.dkr.ecr.eu-central-1.amazonaws.com
```

## K8s

- `aws-auth-cm.yaml` is to give access to github user to access the eks cluster

## Nice to have tools

- https://k9scli.io/: `:resource` eg; :ns (namespace), :pod (pods)

## Steps to deploy (TODO: use TF)

- Create RDS database
- Create ecr repository
- Create github actions role with ecr and secrets manager permissions
- Store secrets in secrets manager
- Load secrets from secrets manager in gh action
- Create eks cluster, node group and nodes (min t3.small - 11 pods)
- Give github user access to cluster with config map (Argo CD in the future)

## NGINX ingress controller for AWS

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.1/deploy/static/provider/aws/deploy.yaml
```

- https://cert-manager.io/docs/usage/ingress/
