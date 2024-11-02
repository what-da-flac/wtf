# Use the official AWS Lambda Python image
FROM public.ecr.aws/lambda/python:3.11
ARG LAMBDA_NAME=""

WORKDIR ${LAMBDA_TASK_ROOT}

# Copy function code to the container
COPY lambdas/${LAMBDA_NAME} .

RUN pip install --upgrade pip && \
    pip install -r requirements.txt

# Command to run the Lambda function
CMD ["app.lambda_handler"]
