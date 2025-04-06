# node-gateway

This is a dummy NodeJS application that routes requests with prefix `/system` to internal Golang Rest API.

## Examples

Using Redirect to another service.

```bash
curl -s http://localhost:3000/system/v1/healthz | jq
```

Response

```json
{
  "ok": true,
  "version": "dev"
}
```

Using internal NodeJS endpoint.

```bash
curl -s http://localhost:3000/hello | jq
```

Response

```json
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe"
}
```