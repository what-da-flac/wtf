FROM golang:1.23 AS builder
ARG SERVICE_VERSION=latest
ARG GITHUB_ACCESS_TOKEN=""
ENV GOPRIVATE=github.com/tech-component/*
WORKDIR /app
COPY . .
RUN echo "machine github.com login ${GITHUB_ACCESS_TOKEN}" > ~/.netrc
RUN cd gateway && make build
FROM public.ecr.aws/docker/library/alpine:3.16
RUN apk update && \
    apk add ca-certificates libc6-compat curl && \
    rm -rf /var/cache/apk/*
WORKDIR /app

COPY --from=builder /app/gateway/service /app/service
ENTRYPOINT ["/app/service"]
