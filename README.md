# wtf

Monorepo for all WTF applications.

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
