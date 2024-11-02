from app import lambda_handler

if __name__ == '__main__':
    event = {
        "id": "abc-123",
        "file": "/Users/mau/Downloads/xxx/sample.torrent",
    }
    res = lambda_handler(event, None)
print(res)
