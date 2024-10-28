ARG BASE_GO_BUILD_IMAGE
ARG BASE_GO_VERSION
FROM ${BASE_GO_BUILD_IMAGE}:${BASE_GO_VERSION} AS builder
ARG GITHUB_ACCESS_TOKEN=""
ARG LAMBDA_NAME=""

WORKDIR /app

COPY . .
RUN echo "machine github.com login ${GITHUB_ACCESS_TOKEN}" > ~/.netrc
RUN CGO_ENABLED=0 go build -buildvcs=false -a --installsuffix cgo -o lambda ./${LAMBDA_NAME}

FROM public.ecr.aws/docker/library/alpine:3.16

RUN apk update && \
    apk add --no-cache ca-certificates libc6-compat aria2 transmission-cli && \
    rm -rf /var/cache/apk/*
WORKDIR /app

COPY --from=builder /app/lambda /app/lambda
ENTRYPOINT ["/app/lambda"]
