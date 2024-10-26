# wtf

Monorepo for all WTF applications.

## Deployments

There are some naming conventions since we are using a monorepo, in order to determine which triggers should be activated during deployments.

`docker.*` starts building all docker images and pushes them to ecr, so other builds can make use of them.

`cdk.*` runs cdk code for all stacks.