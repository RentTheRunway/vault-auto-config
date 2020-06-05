# Global vars
ARG SOURCE=/go/src/cmd/vault-auto-config
ARG BINARY=/bin/vault-auto-config
ARG TEST_SOURCE=/go/src/internal/vault-auto-config

# Build stage
FROM golang:1.14-alpine as build-stage
ARG SOURCE
ARG BINARY
ARG TEST_SOURCE

COPY . /go/src

ENV CGO_ENABLED=0
WORKDIR /go/src
RUN go build -o "${BINARY}" "${SOURCE}"
RUN go test "${TEST_SOURCE}"

# Final image stage
FROM alpine
ARG BINARY

COPY --from=build-stage "${BINARY}" "${BINARY}"

ENV BINARY=${BINARY}
ENTRYPOINT ["/bin/vault-auto-config"]
