import json

from aws_lambda_powertools.utilities.typing import LambdaContext


def lambda_handler(event: dict, context: LambdaContext):
    name = event.get('name', 'stranger')
    res = {
        "greeting": f"Hello, {name}!",
        "statusCode": 200,
    }
    return json.dumps(res)
