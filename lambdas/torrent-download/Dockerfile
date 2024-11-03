ARG BASE_GO_BUILD_IMAGE
ARG BASE_GO_VERSION
FROM ${BASE_GO_BUILD_IMAGE}:${BASE_GO_VERSION} AS builder
ARG GITHUB_ACCESS_TOKEN=""
ARG LAMBDA_NAME=""

WORKDIR /app
RUN pwd && ls -l
COPY . .
RUN echo "machine github.com login ${GITHUB_ACCESS_TOKEN}" > ~/.netrc
RUN cd lambdas && \
    CGO_ENABLED=0 go build -buildvcs=false -a --installsuffix cgo -o lambda ./${LAMBDA_NAME}

FROM public.ecr.aws/docker/library/alpine:3.16

RUN apk update && \
    apk add --no-cache ca-certificates libc6-compat transmission-cli transmission-daemon && \
    rm -rf /var/cache/apk/*
WORKDIR /app

COPY --from=builder /app/lambdas/lambda /app/lambda
ENTRYPOINT ["/app/lambda"]
