ARG GOLANG_VERSION=1.23
FROM public.ecr.aws/docker/library/golang:${GOLANG_VERSION}

ENV LINT_VERSION=v1.61.0
ENV MOQ_VERSION=v0.5.0
ENV NODE_VERSION=20
ENV OPEN_API_VERSION=v2.2.0
ENV REDOCLY_VERSION=1.25.7

RUN apt update && apt install -y unzip
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip > /dev/null && \
    ./aws/install && \
    rm -rf aws
RUN curl -fsSL "https://deb.nodesource.com/setup_${NODE_VERSION}.x" | sh
RUN apt install nodejs -y && \
    npm i -g "@redocly/cli@${REDOCLY_VERSION}" && \
    npm i -g aws-cdk
RUN go install "github.com/matryer/moq@${MOQ_VERSION}" && \
    go install "github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@${OPEN_API_VERSION}" && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin ${LINT_VERSION}

