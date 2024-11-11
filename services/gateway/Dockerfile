ARG BASE_GO_BUILD_IMAGE
ARG BASE_GO_VERSION
FROM ${BASE_GO_BUILD_IMAGE}:${BASE_GO_VERSION} AS builder
ARG SERVICE_VERSION=latest
ARG SERVICE_NAME=""

WORKDIR /app
COPY . .

RUN cd services/${SERVICE_NAME} && go install

FROM public.ecr.aws/docker/library/alpine:3.16
ARG SERVICE_NAME=""
RUN apk update && \
    apk add ca-certificates libc6-compat curl && \
    rm -rf /var/cache/apk/*
WORKDIR /app

COPY --from=builder /go/bin/${SERVICE_NAME} /app/service
ENTRYPOINT ["/app/service"]
