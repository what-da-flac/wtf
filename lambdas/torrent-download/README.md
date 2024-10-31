# torrent-download

## Run Locally

You can run this lambda locally without any other no sense tooling.

### PyCharm

```bash
python local.py
```

### Docker

Run `make docker-build-all`. Once all lambda images have been built, run this command.

```bash
docker run --rm -p 9000:8080 torrent-download:lambda.0.0.9
```

Wait for docker container to be up and running, then invoke like this:

```bash
curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{"name":"Mauricio"}'
```