ARG BASE_GO_BUILD_IMAGE
ARG BASE_GO_VERSION
FROM ${BASE_GO_BUILD_IMAGE}:${BASE_GO_VERSION} AS builder
ARG SERVICE_VERSION=latest
ARG GITHUB_ACCESS_TOKEN=""
ARG SERVICE_NAME=""
ENV GOPRIVATE=github.com/what-da-flac/*

WORKDIR /app
COPY . .

RUN echo "machine github.com login ${GITHUB_ACCESS_TOKEN}" > ~/.netrc
RUN cd services/${SERVICE_NAME} && make build

FROM public.ecr.aws/docker/library/alpine:3.16
ARG SERVICE_NAME=""
RUN apk update && \
    apk add ca-certificates libc6-compat aria2 transmission-cli && \
    rm -rf /var/cache/apk/*
WORKDIR /app

COPY --from=builder /app/services/${SERVICE_NAME}/service /app/service
ENTRYPOINT ["/app/service"]
