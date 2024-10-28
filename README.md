# wtf

Monorepo for all WTF applications.

## Deployments

All deployments start when you push a git tag. Based on the tag name, actions will be triggered automatically.

There are some naming conventions since we are using a monorepo, in order to determine which triggers should be activated during deployments.

`docker.*` Deploys docker images used among builds.

`cdk.*` Deploys cdk stacks.

`gateway.*` Deploys `gateway` service.

`lambda.*` Deploys `lambdas` services.

`ui.*` Deploys `ui` application to S3.
