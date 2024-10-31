import json


def lambda_handler(event: dict, context):
    name = event.get('name', 'stranger')
    return {
        "statusCode": 200,
        "body": json.dumps(f"Hello, {name}!")
    }
