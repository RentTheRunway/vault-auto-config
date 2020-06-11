# Global vars
ARG SOURCE=/go/src/cmd/vault-auto-config
ARG BINARY=/bin/vault-auto-config

# Build stage
FROM golang:1.14-alpine as build-stage
ARG SOURCE
ARG BINARY

COPY . /go/src

ENV CGO_ENABLED=0
WORKDIR /go/src
RUN go build -o "${BINARY}" "${SOURCE}"

# Final image stage
FROM alpine
ARG BINARY

COPY --from=build-stage "${BINARY}" "${BINARY}"

ENV BINARY=${BINARY}
ENTRYPOINT ["/bin/vault-auto-config"]
