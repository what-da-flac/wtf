ARG BASE_GO_BUILD_IMAGE
ARG BASE_GO_VERSION
FROM ${BASE_GO_BUILD_IMAGE}:${BASE_GO_VERSION} AS builder
ARG SERVICE_VERSION=latest
ARG SERVICE_NAME=""
ARG GIT_TAG=""

WORKDIR /app
COPY . .

RUN cd services/${SERVICE_NAME} && \
    go build -o service  -ldflags="-X 'main.Version=${GIT_TAG}'"

FROM public.ecr.aws/docker/library/alpine:3.16
ARG SERVICE_NAME=""
RUN apk update && \
    apk add ca-certificates libc6-compat aria2 && \
    rm -rf /var/cache/apk/*
WORKDIR /app

COPY --from=builder /app/services/${SERVICE_NAME}/service /app/service
ENTRYPOINT ["/app/service"]
