from app import lambda_handler

if __name__ == '__main__':
    event = {
        "name": "Mauricio",
    }
    res = lambda_handler(event, None)
    print(res)
