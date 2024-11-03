# torrent-download

This lambda receives SQS message which contains a torrent as payload.

Decided to move on with Go, because getting errors on stupid tasks such as initializing/locking pipenv.

```json
{
  "id": "77be913c-6cbe-4db2-bba0-0ae6e755f63f",
  "filename": "my-download.torrent"
}
```

## Requirements

This lambda leverages `transmission-cli` programs, so a transmission-daemon 
needs to be running for this lambda to work properly.

```bash
transmission-daemon --foreground
```

## Run Locally

You can run this lambda locally without any other no sense tooling.

### PyCharm

```bash
python local.py
```

### Docker

Run `make build-all`. Once all lambda images have been built, run this command.

```bash
docker run --rm -p 9000:8080 torrent-download:lambda.0.0.9
```

Wait for docker container to be up and running, then invoke like this:

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":"Mauricio"}'
```