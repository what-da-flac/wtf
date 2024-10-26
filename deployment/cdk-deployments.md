# cdk

This repository stores all infrastructure as code to be used on any AWS account.

CDK deployments are done manually for the time being. All the other assets are automatically updated.

## Docker Images

Some docker images are stored in this repository and they should be updated using a codebuild job.

Codebuild script lives at [docker.yaml](./docker.yaml).

Since it would collide with regular deployments, a Codebuild job should be created manually to handle it. These are the
parameters to such job:

Github Repository: https://github.com/tech-component/wtf-deployment.git

Source Version: empty

Rebuild every time a code change is pushed to this repository: unchecked

Build Type: Single

Filter Group: None

Provisioned Model: On demand

Environment Image: Managed Image

Compute: EC2

Operating System: Ubuntu

Runtime: Standard

Image: aws/codebuild/standard:7.0

Image Version: use latest always

Timeout: 1 hour

Queued timeout: 8 hours

Privileged: checked

Certificate: None

VPC: None

Compute: 3 Gb

Environment Variables

| Name            | Type   | Value                       |
|-----------------|--------|-----------------------------|
| GO_PRIVATE      | text   | github.com/what-da-flac/* |
| SSH_PRIVATE_KEY | secret | mauleyzaola-private-key     |

Use a buildspec file: `docker.yaml`

Artifacts: None

CloudWatch Logs: Leave default `aws/codebuild/docker`

Service Role: Adjust as needed considering the least access privilege. This is a working example

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "AllowECRActions",
      "Effect": "Allow",
      "Action": [
        "ecr:GetAuthorizationToken"
      ],
      "Resource": "*"
    },
    {
      "Sid": "AllowECRPush",
      "Effect": "Allow",
      "Action": [
        "ecr:InitiateLayerUpload",
        "ecr:UploadLayerPart",
        "ecr:CompleteLayerUpload",
        "ecr:PutImage",
        "ecr:BatchCheckLayerAvailability"
      ],
      "Resource": "arn:aws:ecr:us-east-2:160885250498:repository/wtf-go-builder"
    },
    {
      "Sid": "AllowSecretsManagerAccess",
      "Effect": "Allow",
      "Action": [
        "secretsmanager:GetSecretValue"
      ],
      "Resource": "arn:aws:secretsmanager:us-east-2:160885250498:secret:*"
    }
  ]
}
```

## File Structure

*`docker.yaml`* This is a codebuild script that should be triggered manually since it is not listening to any event,
like the others. Its goal is to update docker images we store in ECR to build other things.

*`configs/`* This directory stores all the parameters to CDK as YAML files. Golang structs are mapped against AWS
objects using the same name in most cases. Occasionally, there is little magic to use enums, that's all.

*`main.go`* basically just parses YAML files into `go-common`
cdk [types](https://github.com/what-da-flac/wtf/go-common/blob/main/cdk/stacks/stacks.go).

## Conventions

Some functionality is prepared and fixed in the code, to stop thinking too much about the parameters, amoung these are
resources to determine operational costs and others performance related.

## Execution

Get a difference of the stack

```
make diff
```

Deploy changes

```
make deploy
```

Destroy stack

```
make destroy
```

## Secrets

In order to make this work, we need to store these secrets:

`gh-mauleyzaola-token`

```json
{
  "ServerType": "GITHUB",
  "AuthType": "PERSONAL_ACCESS_TOKEN",
  "Token": "ghp_xxx"
}
```

`gh-token`

```
ghp_xxx
```

`mauleyzaola-private-key`

```shell
-----BEGIN OPENSSH PRIVATE KEY-----
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx==
-----END OPENSSH PRIVATE KEY-----
```

`google-api-key`

```
xxxxx
```

`google-client-secret`

```
xxx
```

`google-client-id`

```
xxx
```