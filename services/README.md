# services

This directory contains all backend services. File structure supports services on any programming language. A Dockerfile image definition should be present for each service.

Since all services are written in Golang, no specific considerations are for the time being (like unit tests, mocks, etc). If additional programming languages are used in the future, we should adapt Makefile/Codebuild accordingly.

## Local Workflow

Local is considered for local development, and services are running from and IDE or directly on the terminal.

Docker compose only runs the dependent services, such as databases and message brokers.

Start

```bash
make local-start
```

Logs

```bash
make local-logs
```

Stop

```bash
make local-stop
```

## Dockerized Execution

Consider UI application is already running. First need to build services:

```bash
make build-all
```

Then start services

```bash
make docker-start
```

To stop services, run this command

```bash
make docker-stop
```
